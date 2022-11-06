# Lockfer

## Protect your files, protect your privacy.

Lockfer allows you to share confidential or sensitive files. Thanks to its encryption system, your data is no longer readable until your correspondent accesses it.

Upload your files allows you to get a sharing link. This link give access to your encrypted files.


## Back-end

An api based on Go allows you to encrypt your files. Files are saved on your server.

```sh
Go version: 1.19
Modules used:
- httprate
- jwt
- mux
- godotenv
- gouuid
- crypto
- time
- xxhash
```

### Specs

```sh
/api/v1/upload: upload and encrypt your files
/api/v1/decrypt: decrypt files
/api/v1/download/{{uuid}}: download your archive with decrypted files
```


## Front-end

This project was generated with [Angular CLI](https://github.com/angular/angular-cli) version 13.3.9.

### Development server

Run `ng serve` for a dev server. Navigate to `http://localhost:4200/`. The application will automatically reload if you change any of the source files.

### Code scaffolding

Run `ng generate component component-name` to generate a new component. You can also use `ng generate directive|pipe|service|class|guard|interface|enum|module`.

### Build

Run `ng build` to build the project. The build artifacts will be stored in the `dist/` directory.

### Running unit tests

Run `ng test` to execute the unit tests via [Karma](https://karma-runner.github.io).

### Running end-to-end tests

Run `ng e2e` to execute the end-to-end tests via a platform of your choice. To use this command, you need to first add a package that implements end-to-end testing capabilities.

### Further help

To get more help on the Angular CLI use `ng help` or go check out the [Angular CLI Overview and Command Reference](https://angular.io/cli) page.
