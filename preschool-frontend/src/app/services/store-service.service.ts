import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class StoreServiceService {

  constructor() { }

  getRoleFromToken(): string {
    let token = window.localStorage.getItem('authToken')
    if (token) {
      let tokenSplit = token.split('.')
      let decoded = decodeURIComponent(encodeURIComponent(window.atob(tokenSplit[1])))
      let obj = JSON.parse(decoded)
      return obj['userType']
    }
    return ""
  }

  getJMBGFromToken(): String{
    let token = window.localStorage.getItem('authToken')
    if (token) {
      let tokenSplit = token.split('.')
      let decoded = decodeURIComponent(encodeURIComponent(window.atob(tokenSplit[1])))
      let obj = JSON.parse(decoded)
      return obj['jmbg']
    }
    return ""
  }

  getServiceFromStore(): any {
     return localStorage.getItem('service')
  }

  getCustomRole(): any {
    return localStorage.getItem('customRole')
  }

}
