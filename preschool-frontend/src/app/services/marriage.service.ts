import { Injectable } from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {Observable} from "rxjs";
import {Marriage} from "../models/marriage";
import * as http from "http";
import {environment} from "../../environments/environment";
import {StoreServiceService} from "./store-service.service";
import { UserDied } from '../models/userDied.model';

@Injectable({
  providedIn: 'root'
})
export class MarriageService {

  private url = "registrar";

  constructor(
    private http: HttpClient,
    private storeService: StoreServiceService
  ) { }

  public CreateMarriage(marriage: Marriage): Observable<any> {
    return this.http.post(`${environment.baseApiUrl}/${this.url}/marriage`, marriage);
  }

  public GetCertificate(type: String): Observable<any> {
    return this.http.get(`${environment.baseApiUrl}/${this.url}/certificate/` + this.storeService.getJMBGFromToken() + '/' + type);
  }

  public UpdateCertificate(userDied: UserDied): Observable<string> {
    return this.http.post<string>(`${environment.baseApiUrl}/${this.url}/died`, userDied);
  }
}
