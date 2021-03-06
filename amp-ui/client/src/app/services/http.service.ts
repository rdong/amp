import { Injectable } from '@angular/core';
import { Http, Headers, Response } from '@angular/http';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/catch';
import { Subject } from 'rxjs/Subject';
import {Observable} from 'rxjs/Observable';
import { User } from '../models/user.model';
import { Stack } from '../models/stack.model';
import { Endpoint } from '../models/endpoint.model';

@Injectable()
export class HttpService {
  private token = ""
  private currentEndpointName = ""
  onHttpError = new Subject();
  constructor(private http : Http) {}

  setToken(endpointName: string, token : string) {
    this.currentEndpointName = endpointName
    this.token=token
  }

  private setHeaders() {
    var headers = new Headers
    headers.set('TokenKey', this.token)
    headers.set('Endpoint', this.currentEndpointName)
    return headers
  }

  endpoints() {
    return this.http.get("/api/v1/endpoints")
      .map((res : Response) => res.json())
  }

  connectEndpoint(endpoint : Endpoint) {
    return this.http.post("/api/v1/connect", endpoint);
  }

  users() {
    return this.http.get("/api/v1/users", { headers: this.setHeaders() })
      .map((res : Response) => {
        const data = res.json()
        let list : User[] = []
        for (let item of data.users) {
          let user = new User(item.name, item.email, "User")
          user.verified = item.is_verified
          list.push(user)
        }
        return list
      })
  }

  login(user : User, pwd : string) {
    return this.http.post("/api/v1/login", {name: user.name, pwd: pwd}, { headers: this.setHeaders() });
  }

  stacks() {
    return this.http.get("/api/v1/stacks", { headers: this.setHeaders() })
    .map((res : Response) => {
      const data = res.json()
      let list : Stack[] = []
      for (let item of data.stacks) {
        let stack = new Stack(
          item.stack.id,
          item.stack.name,
          item.service,
          item.stack.owner.name,
          item.stack.owner.type
        )
        list.push(stack)
      }
      return list
    })
  }

}
