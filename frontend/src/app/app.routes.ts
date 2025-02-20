import { Routes } from '@angular/router';
import {AuthGuard} from './authorization/auth-guard.service';

export const routes: Routes = [
  {
    path: "",
    pathMatch: "full",
    redirectTo: "candidate",
  },
  {
    path: "candidate",
    loadComponent: () => import('./senior-candidate/senior-candidate.component').then(m => m.SeniorCandidateComponent),
    canActivate: [AuthGuard],
  },
  {
    path: "hr",
    loadComponent: () => import('./hr/hr.component').then(m => m.HrComponent),
    canActivate: [AuthGuard],
  },
  {
    path: "management",
    loadComponent: () => import('./management-dashboard/management-dashboard.component').then(m => m.ManagementDashboardComponent),
    canActivate: [AuthGuard],
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
    path: "unauthorized",
    loadComponent: () => import('./authorization/unauthorized.component').then(m => m.UnauthorizedComponent),
  },
  {
    path: "**",
    redirectTo: "",
  },
];
