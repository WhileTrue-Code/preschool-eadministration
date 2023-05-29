import { Injectable } from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {Observable} from "rxjs";
import {Marriage} from "../models/marriage";
import * as http from "http";
import {environment} from "../../environments/environment";

@Injectable({
  providedIn: 'root'
})
export class MarriageService {

  private url = "registrar";

  constructor(
    private http: HttpClient
  ) { }

  public CreateMarriage(marriage: Marriage): Observable<any> {
    return this.http.post(`${environment.baseApiUrl}/${this.url}/marriage`, marriage);
  }
}
