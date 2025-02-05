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
    path: "admin",
    loadComponent: () => import('./admin/admin.component') .then(m => m.AdminComponent),
    canActivate: [AuthGuard],
  },
  {
    path: "employer",
    loadComponent: () => import('./employer/employer.component').then(m => m.EmployerComponent),
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
