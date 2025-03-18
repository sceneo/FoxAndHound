import { Component, OnInit } from "@angular/core";
import { environment } from "../../environments/environment";
import { Router } from "@angular/router";

const DEFAULT_ROUTE = '/candidate';
const ERROR_ROUTE = '/unauthorized';

@Component({
    selector: 'app-auth',
    standalone: true,
    template: `<p>Processing authentication...</p>`,
})
export class AuthComponent implements OnInit {
    constructor(private router: Router) {}

    ngOnInit(): void {
        if (!environment.useMsalAuth) {
            this.router.navigate(['/candidate']);
        }
    }
}