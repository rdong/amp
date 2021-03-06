package account

import (
	"log"

	"github.com/appcelerator/amp/api/auth"
	"github.com/appcelerator/amp/cmd/amplifier/server/configuration"
	"github.com/appcelerator/amp/data/accounts"
	"github.com/appcelerator/amp/pkg/mail"
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Server is used to implement account.AccountServer
type Server struct {
	Accounts accounts.Interface
	Mailer   *mail.Mailer
	Config   *configuration.Configuration
}

func convertError(err error) error {
	switch err {
	case accounts.InvalidName:
	case accounts.InvalidEmail:
	case accounts.InvalidToken:
	case accounts.PasswordTooWeak:
		return status.Errorf(codes.InvalidArgument, err.Error())
	case accounts.WrongPassword:
		return status.Errorf(codes.Unauthenticated, err.Error())
	case accounts.UserNotVerified:
	case accounts.AtLeastOneOwner:
	case accounts.TokenAlreadyUsed:
		return status.Errorf(codes.FailedPrecondition, err.Error())
	case accounts.UserAlreadyExists:
	case accounts.EmailAlreadyUsed:
	case accounts.OrganizationAlreadyExists:
	case accounts.TeamAlreadyExists:
		return status.Errorf(codes.AlreadyExists, err.Error())
	case accounts.UserNotFound:
	case accounts.OrganizationNotFound:
	case accounts.TeamNotFound:
		return status.Errorf(codes.NotFound, err.Error())
	case accounts.NotAuthorized:
		return status.Errorf(codes.PermissionDenied, err.Error())
	}
	return status.Errorf(codes.Internal, err.Error())
}

func getServerAddress(ctx context.Context) string {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return ""
	}
	authorities := md[":authority"]
	if len(authorities) == 0 {
		return ""
	}
	return authorities[0]
}

func (s *Server) getRequesterEmail(ctx context.Context) string {
	activeOrganization := auth.GetActiveOrganization(ctx)
	if activeOrganization != "" {
		organization, err := s.Accounts.GetOrganization(ctx, activeOrganization)
		if err != nil {
			return ""
		}
		if organization == nil {
			return ""
		}
		return organization.Email
	}

	user, err := s.Accounts.GetUser(ctx, auth.GetUser(ctx))
	if err != nil {
		return ""
	}
	if user == nil {
		return ""
	}
	return user.Email
}

// Users

// SignUp implements account.SignUp
func (s *Server) SignUp(ctx context.Context, in *SignUpRequest) (*empty.Empty, error) {
	// Create user
	user, err := s.Accounts.CreateUser(ctx, in.Name, in.Email, in.Password)
	if err != nil {
		s.Accounts.DeleteUser(ctx, in.Name)
		return nil, convertError(err)
	}

	switch s.Config.Registration {
	case configuration.RegistrationEmail:
		// Create a verification token valid for an hour
		token, err := auth.CreateVerificationToken(user.Name)
		if err != nil {
			s.Accounts.DeleteUser(ctx, in.Name)
			return nil, convertError(err)
		}

		// Send the verification email
		if err := s.Mailer.SendAccountVerificationEmail(user.Email, user.Name, token, getServerAddress(ctx)); err != nil {
			s.Accounts.DeleteUser(ctx, in.Name)
			return nil, convertError(err)
		}
	}

	log.Println("Successfully created user", user.Name)
	return &empty.Empty{}, nil

}

// Verify implements account.Verify
func (s *Server) Verify(ctx context.Context, in *VerificationRequest) (*empty.Empty, error) {
	// Validate the token
	claims, err := auth.ValidateToken(in.Token, auth.TokenTypeVerification)
	if err != nil {
		return nil, accounts.InvalidToken
	}
	if err := s.Accounts.VerifyUser(ctx, claims.AccountName); err != nil {
		return nil, convertError(err)
	}
	user, err := s.Accounts.GetUser(ctx, claims.AccountName)
	if err != nil {
		return nil, convertError(err)
	}
	if err := s.Mailer.SendAccountVerifiedEmail(user.Email, user.Name); err != nil {
		return nil, convertError(err)
	}
	log.Println("Successfully verified user", user.Name)
	return &empty.Empty{}, nil
}

// Login implements account.Login
func (s *Server) Login(ctx context.Context, in *LogInRequest) (*empty.Empty, error) {
	// Check password
	if err := s.Accounts.CheckUserPassword(ctx, in.Name, in.Password); err != nil {
		return nil, convertError(err)
	}
	// Create an authentication token valid for a day
	token, err := auth.CreateLoginToken(in.Name, "")
	if err != nil {
		return nil, convertError(err)
	}
	// Send the auth token to the client
	md := metadata.Pairs(auth.TokenKey, token)
	if err := grpc.SendHeader(ctx, md); err != nil {
		return nil, convertError(err)
	}
	log.Println("Successfully logged user in", in.Name)
	return &empty.Empty{}, nil
}

// PasswordReset implements account.PasswordReset
func (s *Server) PasswordReset(ctx context.Context, in *PasswordResetRequest) (*empty.Empty, error) {
	// Get the user
	user, err := s.Accounts.GetUser(ctx, in.Name)
	if err != nil {
		return nil, convertError(err)
	}
	if user == nil {
		return nil, status.Errorf(codes.NotFound, "user not found: %s", in.Name)
	}
	// Create a password reset token valid for an hour
	token, err := auth.CreatePasswordToken(user.Name)
	if err != nil {
		return nil, convertError(err)
	}
	// Send the password reset email
	if err := s.Mailer.SendAccountResetPasswordEmail(user.Email, user.Name, token, getServerAddress(ctx)); err != nil {
		return nil, convertError(err)
	}
	log.Println("Successfully reset password for user", user.Name)
	return &empty.Empty{}, nil
}

// PasswordSet implements account.PasswordSet
func (s *Server) PasswordSet(ctx context.Context, in *PasswordSetRequest) (*empty.Empty, error) {
	// Validate token
	claims, err := auth.ValidateToken(in.Token, auth.TokenTypePassword)
	if err != nil {
		return nil, convertError(err)
	}
	// Sets the new password
	if err := s.Accounts.SetUserPassword(ctx, claims.AccountName, in.Password); err != nil {
		return &empty.Empty{}, convertError(err)
	}
	log.Println("Successfully set new password for user", claims.AccountName)
	return &empty.Empty{}, nil
}

// PasswordChange implements account.PasswordChange
func (s *Server) PasswordChange(ctx context.Context, in *PasswordChangeRequest) (*empty.Empty, error) {
	// Get requesting user
	requester := auth.GetUser(ctx)

	// Check the existing password
	if err := s.Accounts.CheckUserPassword(ctx, requester, in.ExistingPassword); err != nil {
		return nil, convertError(err)
	}
	// Sets the new password
	if err := s.Accounts.SetUserPassword(ctx, requester, in.NewPassword); err != nil {
		return &empty.Empty{}, convertError(err)
	}
	log.Println("Successfully updated the password for user", requester)
	return &empty.Empty{}, nil
}

// ForgotLogin implements account.PasswordChange
func (s *Server) ForgotLogin(ctx context.Context, in *ForgotLoginRequest) (*empty.Empty, error) {
	// Get the user
	user, err := s.Accounts.GetUserByEmail(ctx, in.Email)
	if err != nil {
		return nil, convertError(err)
	}
	if user == nil {
		return nil, status.Errorf(codes.NotFound, "user not found: %s", in.Email)
	}
	// Send the account name reminder email
	if err := s.Mailer.SendAccountNameReminderEmail(user.Email, user.Name); err != nil {
		return nil, convertError(err)
	}
	log.Println("Successfully processed forgot login request for user", user.Name)
	return &empty.Empty{}, nil
}

// GetUser implements account.GetUser
func (s *Server) GetUser(ctx context.Context, in *GetUserRequest) (*GetUserReply, error) {
	// Get the user
	user, err := s.Accounts.GetUser(ctx, in.Name)
	if err != nil {
		return nil, convertError(err)
	}
	if user == nil {
		return nil, status.Errorf(codes.NotFound, "user not found: %s", in.Name)
	}
	log.Println("Successfully retrieved user", user.Name)
	return &GetUserReply{User: user}, nil
}

// GetUserOrganizations implements account.GetUserOrganizations
func (s *Server) GetUserOrganizations(ctx context.Context, in *GetUserOrganizationsRequest) (*GetUserOrganizationsReply, error) {
	// Get the user's organizations
	organizations, err := s.Accounts.GetUserOrganizations(ctx, in.Name)
	if err != nil {
		return nil, convertError(err)
	}
	log.Printf("Successfully retrieved organizations for user %s \n", in.Name)
	return &GetUserOrganizationsReply{Organizations: organizations}, nil
}

// ListUsers implements account.ListUsers
func (s *Server) ListUsers(ctx context.Context, in *ListUsersRequest) (*ListUsersReply, error) {
	users, err := s.Accounts.ListUsers(ctx)
	if err != nil {
		return nil, convertError(err)
	}
	log.Println("Successfully list users")
	return &ListUsersReply{Users: users}, nil
}

// DeleteUser implements account.DeleteUser
func (s *Server) DeleteUser(ctx context.Context, in *DeleteUserRequest) (*empty.Empty, error) {
	// Get requesting user
	user, err := s.Accounts.GetUser(ctx, in.Name)
	if err != nil {
		return nil, convertError(err)
	}
	if user == nil {
		return nil, status.Errorf(codes.NotFound, "user not found: %s", in.Name)
	}

	if err := s.Accounts.DeleteUser(ctx, in.Name); err != nil {
		return nil, convertError(err)
	}
	if err := s.Mailer.SendAccountRemovedEmail(user.Email, user.Name); err != nil {
		return nil, convertError(err)
	}
	log.Println("Successfully deleted user", in.Name)
	return &empty.Empty{}, nil
}

// Switch implements account.Switch
func (s *Server) Switch(ctx context.Context, in *SwitchRequest) (*empty.Empty, error) {
	// Get user name
	userName := auth.GetUser(ctx)

	activeOrganization := ""
	// If the account name is not his own account, it has to be an organization
	if userName != in.Account {
		organization, err := s.Accounts.GetOrganization(ctx, in.Account)
		if err != nil {
			return nil, convertError(err)
		}
		if organization == nil {
			return nil, status.Errorf(codes.NotFound, "organization not found: %s", in.Account)
		}
		if !organization.HasMember(userName) {
			return nil, status.Errorf(codes.FailedPrecondition, "user %s is not a member of organization  %s", userName, in.Account)
		}
		activeOrganization = organization.Name
	}

	// Create an authentication token valid for a day
	token, err := auth.CreateLoginToken(userName, activeOrganization)
	if err != nil {
		return nil, convertError(err)
	}
	// Send the auth token to the client
	md := metadata.Pairs(auth.TokenKey, token)
	if err := grpc.SendHeader(ctx, md); err != nil {
		return nil, convertError(err)
	}
	return &empty.Empty{}, nil
}

// Organizations

// CreateOrganization implements account.CreateOrganization
func (s *Server) CreateOrganization(ctx context.Context, in *CreateOrganizationRequest) (*empty.Empty, error) {
	if err := s.Accounts.CreateOrganization(ctx, in.Name, in.Email); err != nil {
		return nil, convertError(err)
	}
	// Send confirmation email
	if email := s.getRequesterEmail(ctx); email != "" {
		if err := s.Mailer.SendOrganizationCreatedEmail(email, in.Name); err != nil {
			return nil, convertError(err)
		}
	}
	log.Println("Successfully created organization", in.Name)
	return &empty.Empty{}, nil
}

// AddUserToOrganization implements account.AddOrganizationMember
func (s *Server) AddUserToOrganization(ctx context.Context, in *AddUserToOrganizationRequest) (*empty.Empty, error) {
	if err := s.Accounts.AddUserToOrganization(ctx, in.OrganizationName, in.UserName); err != nil {
		return &empty.Empty{}, convertError(err)
	}
	// Send confirmation email
	if email := s.getRequesterEmail(ctx); email != "" {
		if err := s.Mailer.SendUserAddedInOrganizationEmail(email, in.OrganizationName, in.UserName); err != nil {
			return nil, convertError(err)
		}
	}
	log.Printf("Successfully added member %s to organization %s\n", in.UserName, in.OrganizationName)
	return &empty.Empty{}, nil
}

// RemoveUserFromOrganization implements account.RemoveOrganizationMember
func (s *Server) RemoveUserFromOrganization(ctx context.Context, in *RemoveUserFromOrganizationRequest) (*empty.Empty, error) {
	if err := s.Accounts.RemoveUserFromOrganization(ctx, in.OrganizationName, in.UserName); err != nil {
		return &empty.Empty{}, convertError(err)
	}
	// Send confirmation email
	if email := s.getRequesterEmail(ctx); email != "" {
		if err := s.Mailer.SendUserRemovedFromOrganizationEmail(email, in.OrganizationName, in.UserName); err != nil {
			return nil, convertError(err)
		}
	}
	log.Printf("Successfully removed user %s from organization %s\n", in.UserName, in.OrganizationName)
	return &empty.Empty{}, nil
}

// ChangeOrganizationMemberRole implements account.ChangeOrganizationMemberRole
func (s *Server) ChangeOrganizationMemberRole(ctx context.Context, in *ChangeOrganizationMemberRoleRequest) (*empty.Empty, error) {
	if err := s.Accounts.ChangeOrganizationMemberRole(ctx, in.OrganizationName, in.UserName, in.Role); err != nil {
		return &empty.Empty{}, convertError(err)
	}
	log.Printf("Successfully changed role of user %s from organization %s to %s\n", in.UserName, in.OrganizationName, in.Role.String())
	return &empty.Empty{}, nil
}

// GetOrganization implements account.GetOrganization
func (s *Server) GetOrganization(ctx context.Context, in *GetOrganizationRequest) (*GetOrganizationReply, error) {
	organization, err := s.Accounts.GetOrganization(ctx, in.Name)
	if err != nil {
		return nil, convertError(err)
	}
	if organization == nil {
		return nil, status.Errorf(codes.NotFound, "organization not found: %s", in.Name)
	}
	log.Println("Successfully retrieved organization", organization.Name)
	return &GetOrganizationReply{Organization: organization}, nil
}

// ListOrganizations implements account.ListOrganizations
func (s *Server) ListOrganizations(ctx context.Context, in *ListOrganizationsRequest) (*ListOrganizationsReply, error) {
	organizations, err := s.Accounts.ListOrganizations(ctx)
	if err != nil {
		return nil, convertError(err)
	}
	log.Println("Successfully list organizations")
	return &ListOrganizationsReply{Organizations: organizations}, nil
}

// DeleteOrganization implements account.DeleteOrganization
func (s *Server) DeleteOrganization(ctx context.Context, in *DeleteOrganizationRequest) (*empty.Empty, error) {
	if err := s.Accounts.DeleteOrganization(ctx, in.Name); err != nil {
		return nil, convertError(err)
	}
	// Send confirmation email
	if email := s.getRequesterEmail(ctx); email != "" {
		if err := s.Mailer.SendOrganizationRemovedEmail(email, in.Name); err != nil {
			return nil, convertError(err)
		}
	}
	log.Println("Successfully deleted organization", in.Name)
	return &empty.Empty{}, nil
}

// Teams

// CreateTeam implements account.CreateTeam
func (s *Server) CreateTeam(ctx context.Context, in *CreateTeamRequest) (*empty.Empty, error) {
	if err := s.Accounts.CreateTeam(ctx, in.OrganizationName, in.TeamName); err != nil {
		return nil, convertError(err)
	}
	// Send confirmation email
	if email := s.getRequesterEmail(ctx); email != "" {
		if err := s.Mailer.SendTeamCreatedEmail(email, in.TeamName); err != nil {
			return nil, convertError(err)
		}
	}
	log.Printf("Successfully created team %s in organization %s\n", in.TeamName, in.OrganizationName)
	return &empty.Empty{}, nil
}

// AddUserToTeam implements account.AddUserToTeam
func (s *Server) AddUserToTeam(ctx context.Context, in *AddUserToTeamRequest) (*empty.Empty, error) {
	if err := s.Accounts.AddUserToTeam(ctx, in.OrganizationName, in.TeamName, in.UserName); err != nil {
		return &empty.Empty{}, convertError(err)
	}
	// Send confirmation email
	if email := s.getRequesterEmail(ctx); email != "" {
		if err := s.Mailer.SendUserAddedInTeamEmail(email, in.TeamName, in.UserName); err != nil {
			return nil, convertError(err)
		}
	}
	log.Printf("Successfully added user %s to team %s in organization %s\n", in.UserName, in.TeamName, in.OrganizationName)
	return &empty.Empty{}, nil
}

// RemoveUserFromTeam implements account.RemoveUserFromTeam
func (s *Server) RemoveUserFromTeam(ctx context.Context, in *RemoveUserFromTeamRequest) (*empty.Empty, error) {
	if err := s.Accounts.RemoveUserFromTeam(ctx, in.OrganizationName, in.TeamName, in.UserName); err != nil {
		return &empty.Empty{}, convertError(err)
	}
	// Send confirmation email
	if email := s.getRequesterEmail(ctx); email != "" {
		if err := s.Mailer.SendUserRemovedFromTeamEmail(email, in.TeamName, in.UserName); err != nil {
			return nil, convertError(err)
		}
	}
	log.Printf("Successfully removed user %s from teams %s in organization %s\n", in.UserName, in.TeamName, in.OrganizationName)
	return &empty.Empty{}, nil
}

// AddResourceToTeam implements account.AddResourceToTeam
func (s *Server) AddResourceToTeam(ctx context.Context, in *AddResourceToTeamRequest) (*empty.Empty, error) {
	if err := s.Accounts.AddResourceToTeam(ctx, in.OrganizationName, in.TeamName, in.ResourceId); err != nil {
		return &empty.Empty{}, convertError(err)
	}
	log.Printf("Successfully added resource %s to team %s in organization %s\n", in.ResourceId, in.TeamName, in.OrganizationName)
	return &empty.Empty{}, nil
}

// RemoveResourceFromTeam implements account.RemoveResourceFromTeam
func (s *Server) RemoveResourceFromTeam(ctx context.Context, in *RemoveResourceFromTeamRequest) (*empty.Empty, error) {
	if err := s.Accounts.RemoveResourceFromTeam(ctx, in.OrganizationName, in.TeamName, in.ResourceId); err != nil {
		return &empty.Empty{}, convertError(err)
	}
	log.Printf("Successfully removed resource %s from teams %s in organization %s\n", in.ResourceId, in.TeamName, in.OrganizationName)
	return &empty.Empty{}, nil
}

// ChangeTeamResourcePermissionLevel implements account.ChangeOrganizationMemberRole
func (s *Server) ChangeTeamResourcePermissionLevel(ctx context.Context, in *ChangeTeamResourcePermissionLevelRequest) (*empty.Empty, error) {
	if err := s.Accounts.ChangeTeamResourcePermissionLevel(ctx, in.OrganizationName, in.TeamName, in.ResourceId, in.PermissionLevel); err != nil {
		return &empty.Empty{}, convertError(err)
	}
	log.Printf("Successfully changed permission level over resource %s in team %s to %s\n", in.ResourceId, in.TeamName, in.PermissionLevel.String())
	return &empty.Empty{}, nil
}

// GetTeam implements account.GetTeam
func (s *Server) GetTeam(ctx context.Context, in *GetTeamRequest) (*GetTeamReply, error) {
	team, err := s.Accounts.GetTeam(ctx, in.OrganizationName, in.TeamName)
	if err != nil {
		return nil, convertError(err)
	}
	if team == nil {
		return nil, status.Errorf(codes.NotFound, "team not found: %s", in.TeamName)
	}
	log.Println("Successfully retrieved team", team.Name)
	return &GetTeamReply{Team: team}, nil
}

// ListTeams implements account.ListTeams
func (s *Server) ListTeams(ctx context.Context, in *ListTeamsRequest) (*ListTeamsReply, error) {
	teams, err := s.Accounts.ListTeams(ctx, in.OrganizationName)
	if err != nil {
		return nil, err
	}
	return &ListTeamsReply{Teams: teams}, nil
}

// DeleteTeam implements account.DeleteTeam
func (s *Server) DeleteTeam(ctx context.Context, in *DeleteTeamRequest) (*empty.Empty, error) {
	if err := s.Accounts.DeleteTeam(ctx, in.OrganizationName, in.TeamName); err != nil {
		return nil, convertError(err)
	}
	// Send confirmation email
	if email := s.getRequesterEmail(ctx); email != "" {
		if err := s.Mailer.SendTeamRemovedEmail(email, in.TeamName); err != nil {
			return nil, convertError(err)
		}
	}
	log.Printf("Successfully deleted team %s from organization %s\n", in.TeamName, in.OrganizationName)
	return &empty.Empty{}, nil
}
