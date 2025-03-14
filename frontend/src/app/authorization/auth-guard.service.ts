import { Injectable } from '@angular/core';
import { CanActivate, Router } from '@angular/router';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class AuthGuard implements CanActivate {
  constructor(private router: Router) {}

  canActivate(): Observable<boolean> | Promise<boolean> | boolean {
    const isAuthenticated = true; //FIXME

    if (!isAuthenticated) {
      this.router.navigate(['/unauthorized']);
      return false;
    }
    return true;
  }
}
