import { Injectable } from '@angular/core';
import { MsalService, MsalBroadcastService } from '@azure/msal-angular';
import { InteractionStatus, EventMessage, EventType, AuthenticationResult } from '@azure/msal-browser';
import { BehaviorSubject, Observable } from 'rxjs';
import { filter } from 'rxjs/operators';
import { environment } from '../../environments/environment';
import { Router } from '@angular/router';

const DEFAULT_ROUTE = '/candidate';
const ERROR_ROUTE = '/unauthorized';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private usernameSubject = new BehaviorSubject<string>('');
  username$: Observable<string> = this.usernameSubject.asObservable();

  constructor(
    private msalService: MsalService,
    private msalBroadcastService: MsalBroadcastService,
    private router: Router
  ) {
    this.initAuth();

    if (environment.useMsalAuth) {
        this.listenForAuthChanges();
        this.handleRedirectObservable();
    }
  }

  getUsername(): Observable<string> {
    return this.username$;
  }

  login(): void {
    if (environment.useMsalAuth) {
      this.msalService.loginRedirect();
    }
  }

  logout(): void {
    if (environment.useMsalAuth) {
      this.msalService.logoutRedirect();
    }
  }
  
  private initAuth(): void {
    if (environment.useMsalAuth) {
      const account = this.msalService.instance.getAllAccounts()[0];
      this.usernameSubject.next(account ? account.username : '');
    } else {
      this.usernameSubject.next('test.user@prodyna.com');
    }
  }

  private listenForAuthChanges(): void {
    this.msalBroadcastService.inProgress$
      .pipe(filter((status: InteractionStatus) => status === InteractionStatus.None))
      .subscribe(() => {
        const account = this.msalService.instance.getAllAccounts()[0];
        this.usernameSubject.next(account ? account.username : '');
      });

    this.msalBroadcastService.msalSubject$.pipe(
        filter((msg: EventMessage) => msg.eventType === EventType.LOGIN_SUCCESS)
    ).subscribe({
        next: (result: EventMessage) => {
            const authResult = result.payload as AuthenticationResult;
            
            if (authResult && authResult.account) {
            this.msalService.instance.setActiveAccount(authResult.account);
            }
            
            this.router.navigate([DEFAULT_ROUTE]);
        },
        error: (error) => {
            console.error('Error during login:', error);
            this.router.navigate([ERROR_ROUTE]);
        },
    });
  }

  private handleRedirectObservable(): void {
    this.msalService.handleRedirectObservable().subscribe({
      next: () => {},
      error: (error) => {
        console.error('Error processing login redirect:', error);
        this.router.navigate([ERROR_ROUTE]);
      }
    });
  }
}
