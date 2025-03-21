import { Component, OnInit } from '@angular/core';
import { RouterLink, RouterOutlet } from '@angular/router';
import { MatIcon } from '@angular/material/icon';
import { MatIconButton } from '@angular/material/button';
import { MatListItem, MatNavList } from '@angular/material/list';
import { MatSidenav, MatSidenavContainer, MatSidenavContent } from '@angular/material/sidenav';
import { MatToolbar } from '@angular/material/toolbar';
import { HttpClient, HttpClientModule } from '@angular/common/http';
import { ReactiveFormsModule } from '@angular/forms';
import { AuthService } from './auth/auth.service';

@Component({
  selector: 'app-root',
  imports: [
    RouterOutlet, HttpClientModule,
    MatIcon, MatIconButton, MatListItem,
    MatNavList, MatSidenav, MatSidenavContainer,
    ReactiveFormsModule,
    MatSidenavContent, MatToolbar, RouterLink],
  providers:[HttpClient],
  templateUrl: './app.component.html',
  standalone: true,
  styleUrl: './app.component.scss'
})
export class AppComponent implements OnInit {
  menuOpen = false;
  username: string = '';

  constructor(
    private authService: AuthService
  ) {}

  ngOnInit(): void {
    this.authService.getUsername().subscribe(username => {
      this.username = username;
    });
  }

  toggleMenu() {
    this.menuOpen = !this.menuOpen;
  }

  closeMenu() {
    this.menuOpen = false;
  }
}
