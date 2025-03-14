# Webapp

## Build Frontend

```bash
cd frontend
docker build . -f ../infra/docker/build_webapp.Dockerfile -t foxandhound_webapp_local
docker tag foxandhound_webapp_local johannesrosskopp/my_private_repository:foxandhound_webapp_dev_latest
docker push johannesrosskopp/my_private_repository:foxandhound_webapp_dev_latest
```

for a release version also tag and push an image with a version like foxandhound_webapp_dev_1_0_0

## Run backend container with host network

```bash
docker run --init --network="host" -e BACKEND_URL=http://127.0.0.1:8080 foxandhound_webapp_local
```

or when the when local ports are blocked get the local ip (see backend README) and use

```bash
docker run \
    --init \
    -e STAGE=dev \
    -e BACKEND_URL=http://${LOCAL_IP}:8080 \
    -p 8087:80 \
    foxandhound_webapp_local
```

# Angular Stuff

This project was generated using [Angular CLI](https://github.com/angular/angular-cli) version 19.1.4.

## Development server

To start a local development server, run:

```bash
ng serve
```

Once the server is running, open your browser and navigate to `http://localhost:4200/`. The application will automatically reload whenever you modify any of the source files.

## Code scaffolding

Angular CLI includes powerful code scaffolding tools. To generate a new component, run:

```bash
ng generate component component-name
```

For a complete list of available schematics (such as `components`, `directives`, or `pipes`), run:

```bash
ng generate --help
```

## Building

To build the project run:

```bash
ng build
```

This will compile your project and store the build artifacts in the `dist/` directory. By default, the production build optimizes your application for performance and speed.

## Running unit tests

To execute unit tests with the [Karma](https://karma-runner.github.io) test runner, use the following command:

```bash
ng test
```

## Running end-to-end tests

For end-to-end (e2e) testing, run:

```bash
ng e2e
```

Angular CLI does not come with an end-to-end testing framework by default. You can choose one that suits your needs.

## Additional Resources

For more information on using the Angular CLI, including detailed command references, visit the [Angular CLI Overview and Command Reference](https://angular.dev/tools/cli) page.
