import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';

//Services
import { UsersService } from './services/users.service';
import { DockerStacksService } from './services/docker-stacks.service';
import { MenuService } from './services/menu.service';
import { AuthGuard } from './services/auth-guard.service';
import { HttpService } from './services/http.service';
import { OrganizationsService } from './services/organizations.service';
import { DockerServicesService } from './services/docker-services.service';
import { DockerContainersService } from './services/docker-containers.service';
import { SwarmsService } from './services/swarms.service';

//Module
import { AppRoutingModule} from './app-routing.module';
import { DropdownComponent } from './dropdown/dropdown.component';

//Directive
import { DropdownDirective } from './directives/dropdown.directive'

//components
import { AppComponent } from './app.component';
import { SignupComponent } from './auth/signup/signup.component';
import { SigninComponent } from './auth/signin/signin.component';
import { AuthComponent } from './auth/auth/auth.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import { NodesComponent } from './nodes/nodes.component';
import { DockerStacksComponent } from './docker-stacks/docker-stacks.component';
import { PasswordComponent } from './password/password.component';
import { SidebarComponent } from './sidebar/sidebar.component';
import { PageheaderComponent } from './pageheader/pageheader.component';
import { UsersComponent } from './users/users.component';
import { AmpComponent } from './amp/amp.component';
import { SwarmsComponent } from './swarms/swarms.component';
import { LogsComponent } from './logs/logs.component';
import { MetricsComponent } from './metrics/metrics.component';
import { OrganizationsComponent } from './organizations/organizations.component';
import { OrganizationComponent } from './organizations/organization/organization.component';
import { DockerStackDeployComponent } from './docker-stacks/docker-stack-deploy/docker-stack-deploy.component';
import { DockerServicesComponent } from './docker-stacks/docker-services/docker-services.component';
import { DockerContainersComponent } from './docker-stacks/docker-containers/docker-containers.component';

@NgModule({
  declarations: [
    AppComponent,
    SignupComponent,
    SigninComponent,
    AuthComponent,
    DashboardComponent,
    NodesComponent,
    PasswordComponent,
    SidebarComponent,
    PageheaderComponent,
    UsersComponent,
    AmpComponent,
    SwarmsComponent,
    LogsComponent,
    MetricsComponent,
    OrganizationsComponent,
    OrganizationComponent,
    DropdownComponent,
    DropdownDirective,
    DockerStacksComponent,
    DockerStackDeployComponent,
    DockerServicesComponent,
    DockerContainersComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    HttpModule,
    AppRoutingModule
  ],
  providers: [
    DockerStacksService,
    DockerServicesService,
    DockerContainersService,
    UsersService,
    MenuService,
    HttpService,
    OrganizationsService,
    SwarmsService,
    AuthGuard
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
