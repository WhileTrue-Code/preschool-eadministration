import { HttpClient, HttpHeaders } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { environment } from "src/environments/environment";
import { User } from "../models/user.model";

@Injectable({
    providedIn: 'root'
    })
    export class AuthService {
        private url = "auth";
        constructor(private http: HttpClient) { }
    
        public Register(user: User): Observable<string> {
            return this.http.post<string>(`${environment.baseApiUrl}/${this.url}/register`, user);
        }
    
    }