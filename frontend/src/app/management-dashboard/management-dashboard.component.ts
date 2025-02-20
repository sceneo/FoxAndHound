import { Component } from '@angular/core';
import {MatFormField} from '@angular/material/form-field';
import {MatOption, MatSelect} from '@angular/material/select';

@Component({
  selector: 'app-root',
  imports: [
    MatFormField,
    MatSelect,
    MatOption
  ],
  templateUrl: './management-dashboard.component.html',
  standalone: true,
  styleUrl: './management-dashboard.component.scss'
})
export class ManagementDashboardComponent {
}
