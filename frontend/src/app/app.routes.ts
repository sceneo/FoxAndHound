import { Routes } from '@angular/router';
import { MsalGuard } from '@azure/msal-angular';
import { environment } from '../environments/environment';

const authGuard = environment.useMsalAuth ? [MsalGuard] : [];

export const routes: Routes = [
  {
    path: "",
    pathMatch: "full",
    redirectTo: "candidate",
  },
  {
    path: "candidate",
    loadComponent: () => import('./senior-candidate/senior-candidate.component').then(m => m.SeniorCandidateComponent),
    canActivate: authGuard,
  },
  {
    path: "hr",
    loadComponent: () => import('./hr/hr.component').then(m => m.HrComponent),
    canActivate: authGuard,
  },
  {
    path: "management",
    loadComponent: () => import('./management-dashboard/management-dashboard.component').then(m => m.ManagementDashboardComponent),
    canActivate: authGuard,
  },
  {
    path: "error",
    loadComponent: () => import('./error/error.component').then(m => m.ErrorComponent),
  },
  {
    path: "successfull",
    loadComponent: () => import('./successfull/successfull.component').then(m => m.SuccessfullComponent),
  },
  {
    path: "auth",
    loadComponent: () => import('./auth/auth.component').then(m => m.AuthComponent),
  },
  {
    path: "unauthorized",
    loadComponent: () => import('./authorization/unauthorized.component').then(m => m.UnauthorizedComponent),
  },
  {
    path: "**",
    redirectTo: "",
  },
];
