# Go Role User & Permission

## # Install

1. Clone this repository
    ```
    git clone https://github.com/heryfidiawan07/GO-Lang-RESTful-User-Role-Permission.git
    ```

2. Environment
    ```
    Copy .env.example to .env
    ```

3. Set variable .env file
    ```
    APP_NAME="GO RestFull User Role & Permissions"

    APP_HOST="127.0.0.1:8000"
    DATABASE="root:@tcp(127.0.0.1:3306)/go_rest_role_permissions?charset=utf8mb4&parseTime=True&loc=Local"

    JWT_SECRET=

    AUTH_REDIRECT_URL=

    CLIENT_ID_GITHUB=
    CLIENT_SECRETS_GITHUB=

    CLIENT_ID_GOOGLE=
    CLIENT_SECRETS_GOOGLE=
    ```

4. Run Application
    ```
    go run server.go
    ```

5. Open database and show tables
    ```
    table users
    table roles
    table permissions
    table role_permissions
    ```

## # Auth Module

1. Register User
    #### # Postman

    #### Method: POST
    #### Url: <http://localhost:8000/api/v1/auth/register>
    #### Body, raw, json
    ```json
    {
        "name": "Super Admin",
        "username": "superadmin",
        "email": "superadmin@mail.com",
        "password": "12345678"
    }
    ```
    
    #### Response
    ```json
    {
        "data": {
            "Id": "d8388027-49ff-403c-bad2-234f63256a87",
            "Name": "Super Admin",
            "Username": "superadmin",
            "Email": "superadmin@mail.com",
            "Password": "$2a$04$FGCeMC2U2SEFOuSzc2wE7uMJojP8kFPPv28x5NCh8Vgpv/hZwWd7W",
            "SocialId": "",
            "Provider": "",
            "Avatar": "",
            "RoleId": "",
            "CreatedAt": "2022-08-24T16:00:28.685+07:00",
            "UpdatedAt": "2022-08-24T16:00:28.685+07:00",
            "DeletedAt": null
        },
        "message": "Data created successfully & please login",
        "status": true
    }
    ```
2. Login
    #### # Postman
    #### Method: POST
    #### Url: <http://localhost:8000/api/v1/auth/login>
    #### Body, raw, json
    ```json
    {
        "username": "superadmin",
        "password": "12345678"
    }
    ```

    #### Response
    ```json
    {
        "data": {
            "Id": "d8388027-49ff-403c-bad2-234f63256a87",
            "Name": "Super Admin",
            "Username": "superadmin",
            "Email": "superadmin@mail.com",
            "Password": "$2a$04$FGCeMC2U2SEFOuSzc2wE7uMJojP8kFPPv28x5NCh8Vgpv/hZwWd7W",
            "SocialId": "",
            "Provider": "",
            "Avatar": "",
            "RoleId": "",
            "CreatedAt": "2022-08-24T16:00:28.685+07:00",
            "UpdatedAt": "2022-08-24T16:00:28.685+07:00",
            "DeletedAt": null
        },
        "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjE1MDQ5NzIsImlhdCI6MTY2MTMzMjE3MiwicmVmcmVzaF9pZCI6IjY0OWJlMTVjLWEzMGUtNGNhYi05MDg4LTBhZGUwMDg5YzE0ZSJ9.ZCfMSrTXc0fUUWrBULbSYhWMkd_qQaQE42RgIKLSfhs",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjE0MTg1NzIsImlhdCI6MTY2MTMzMjE3MiwidXNlcl9pZCI6ImQ4Mzg4MDI3LTQ5ZmYtNDAzYy1iYWQyLTIzNGY2MzI1NmE4NyJ9.LQ54GgJKSpEEUPzLKj2_ffIEppQK2UCjoRW8atF6FhQ"

    }
    ```

### *In this step is to set super admin*

1. Create Administrator Permissions
    #### # Postman
    #### Method: POST
    #### Authorization: Bearer Token *(token)*
    #### Url: <http://localhost:8000/api/v1/permission>
    #### Body, raw, json
    ```json
    [    
        {
            "parent_menu": "",
            "parent_id": "",
            "name": "role-delete",
            "alias": "Delete Role",
            "url": "",
            "icon": ""
        },
        {
            "parent_menu": "",
            "parent_id": "",
            "name": "user-update",
            "alias": "Edit User",
            "url": "",
            "icon": ""
        },
        {
            "parent_menu": "",
            "parent_id": "",
            "name": "user-delete",
            "alias": "Delete User",
            "url": "",
            "icon": ""
        },
        {
            "parent_menu": "Administrator",
            "parent_id": "",
            "name": "user-index",
            "alias": "User",
            "url": "user",
            "icon": ""
        },
        {
            "parent_menu": "",
            "parent_id": "",
            "name": "role-show",
            "alias": "Show Role",
            "url": "",
            "icon": ""
        },
        {
            "parent_menu": "",
            "parent_id": "",
            "name": "role-update",
            "alias": "Edit Role",
            "url": "",
            "icon": ""
        },
        {
            "parent_menu": "",
            "parent_id": "",
            "name": "user-show",
            "alias": "Show User",
            "url": "",
            "icon": ""
        },
        {
            "parent_menu": "Administrator",
            "parent_id": "",
            "name": "role-index",
            "alias": "Role",
            "url": "role",
            "icon": ""
        },
        {
            "parent_menu": "",
            "parent_id": "",
            "name": "role-store",
            "alias": "Create Role",
            "url": "",
            "icon": ""
        },
        {
            "parent_menu": "",
            "parent_id": "",
            "name": "user-store",
            "alias": "Create User",
            "url": "",
            "icon": ""
        }
    ]
    ```

    #### Response
    ```json
    {
        "data": [],
        "message": "Data created successfully",
        "status": true
    }
    ```

2. Create Role
    #### # Postman
    #### Method: POST
    #### Authorization: Bearer Token *(token)*
    #### Url: <http://localhost:8000/api/v1/permission/role>
    #### Body, raw, json
    ```json
    {
        "name": "Super Admin",
        "permissions": [
            ... see table permissions to set permissions id of Super Admin
            # If superadmin set all permissions id in this post
            # Example
            "08b7994a-0028-4726-bb30-2dff757cba02",
            "0c016190-34a9-482b-a585-0bdef284e2c4",
        ]
    }
    ```

3. Set User role_id on the table User using role id of the Super Admin

## # Get Auth User Information

1. Me
    #### # Postman
    #### Method: GET
    #### Authorization: Bearer Token *(token)*
    #### Url: <http://localhost:8000/api/v1/me>
    #### Response
    ```json
    {
        "data": {
            "Id": "d8388027-49ff-403c-bad2-234f63256a87",
            "Name": "Super Admin",
            "Username": "superadmin",
            "Email": "superadmin@mail.com",
            "Password": "$2a$04$FGCeMC2U2SEFOuSzc2wE7uMJojP8kFPPv28x5NCh8Vgpv/hZwWd7W",
            "SocialId": "",
            "Provider": "",
            "Avatar": "",
            "RoleId": "f7dfeabb-3a5f-415f-8b84-785b87767760",
            "CreatedAt": "2022-08-24T16:00:28.685+07:00",
            "UpdatedAt": "2022-08-24T16:00:28.685+07:00",
            "DeletedAt": null,
            "Role": {
                "Id": "f7dfeabb-3a5f-415f-8b84-785b87767760",
                "Name": "Super Admin",
                "CreatedAt": "2022-08-24T19:10:57+07:00",
                "UpdatedAt": "2022-08-24T19:11:01+07:00",
                "DeletedAt": null,
                "Permissions": [
                    {
                        "Id": "08b7994a-0028-4726-bb30-2dff757cba02",
                        "ParentMenu": "",
                        "ParentId": "8e4eb2f8-0b0e-4738-ba2e-06ba4522fb97",
                        "Name": "role-delete",
                        "Alias": "Delete Role",
                        "Url": "",
                        "Icon": ""
                    },
                    {
                        "Id": "0c016190-34a9-482b-a585-0bdef284e2c4",
                        "ParentMenu": "",
                        "ParentId": "12db6983-e259-4d62-b8a3-7b6f4b578071",
                        "Name": "user-update",
                        "Alias": "Edit User",
                        "Url": "",
                        "Icon": ""
                    },
                    {
                        "Id": "0e53eb66-6153-428c-8a50-8568f3d3cea7",
                        "ParentMenu": "",
                        "ParentId": "12db6983-e259-4d62-b8a3-7b6f4b578071",
                        "Name": "user-delete",
                        "Alias": "Delete User",
                        "Url": "",
                        "Icon": ""
                    },
                    {
                        "Id": "12db6983-e259-4d62-b8a3-7b6f4b578071",
                        "ParentMenu": "Administrator",
                        "ParentId": "",
                        "Name": "user-index",
                        "Alias": "User",
                        "Url": "user",
                        "Icon": ""
                    },
                    {
                        "Id": "2351944a-4d9a-481c-900c-62e9e876393c",
                        "ParentMenu": "",
                        "ParentId": "8e4eb2f8-0b0e-4738-ba2e-06ba4522fb97",
                        "Name": "role-show",
                        "Alias": "Show Role",
                        "Url": "",
                        "Icon": ""
                    },
                    {
                        "Id": "33f8fb89-fdcc-4708-a865-9867df3a5336",
                        "ParentMenu": "",
                        "ParentId": "8e4eb2f8-0b0e-4738-ba2e-06ba4522fb97",
                        "Name": "role-update",
                        "Alias": "Edit Role",
                        "Url": "",
                        "Icon": ""
                    },
                    {
                        "Id": "5ad17591-7219-4a61-95eb-99b2c3666875",
                        "ParentMenu": "",
                        "ParentId": "12db6983-e259-4d62-b8a3-7b6f4b578071",
                        "Name": "user-show",
                        "Alias": "Show User",
                        "Url": "",
                        "Icon": ""
                    },
                    {
                        "Id": "8e4eb2f8-0b0e-4738-ba2e-06ba4522fb97",
                        "ParentMenu": "Administrator",
                        "ParentId": "",
                        "Name": "role-index",
                        "Alias": "Role",
                        "Url": "role",
                        "Icon": ""
                    },
                    {
                        "Id": "a4a38ea5-4978-4eff-be27-4a495fe96ce0",
                        "ParentMenu": "",
                        "ParentId": "8e4eb2f8-0b0e-4738-ba2e-06ba4522fb97",
                        "Name": "role-store",
                        "Alias": "Create Role",
                        "Url": "",
                        "Icon": ""
                    },
                    {
                        "Id": "cdd63ae7-179a-48a8-adfa-63dc59cd2169",
                        "ParentMenu": "",
                        "ParentId": "12db6983-e259-4d62-b8a3-7b6f4b578071",
                        "Name": "user-store",
                        "Alias": "Create User",
                        "Url": "",
                        "Icon": ""
                    }
                ]
            },
            "Actions": [
                "role-delete",
                "user-update",
                "user-delete",
                "user-index",
                "role-show",
                "role-update",
                "user-show",
                "role-index",
                "role-store",
                "user-store"
            ]
        },
        "message": null,
        "status": true
    }
    ```

## # User Module

1. User Index
    #### # Postman
    #### Method: GET
    #### Authorization: Bearer Token *(token)*
    #### Url: <http://localhost:8000/api/v1/user>
    #### Response
    ```json
    {
        "data": [
            {
                "Id": "d8388027-49ff-403c-bad2-234f63256a87",
                "Name": "Super Admin",
                "Username": "superadmin",
                "Email": "superadmin@mail.com",
                "Password": "$2a$04$FGCeMC2U2SEFOuSzc2wE7uMJojP8kFPPv28x5NCh8Vgpv/hZwWd7W",
                "SocialId": "",
                "Provider": "",
                "Avatar": "",
                "RoleId": "f7dfeabb-3a5f-415f-8b84-785b87767760",
                "CreatedAt": "2022-08-24T16:00:28.685+07:00",
                "UpdatedAt": "2022-08-24T16:00:28.685+07:00",
                "DeletedAt": null
            }
        ],
        "message": null,
        "status": true
    }
    ```

2. Create User
    #### # Postman
    #### Method: POST
    #### Authorization: Bearer Token *(token)*
    #### Url: <http://localhost:8000/api/v1/user>
    #### Body, raw, json
    ```json
    {
        "name": "Other",
        "username": "other",
        "email": "other@mail.com",
        "password": "12345678",
        "role_id": "f7dfeabb-3a5f-415f-8b84-785b87767760"
    }
    ```

    #### Response
    ```json
    {
        "data": {
            "Id": "73907481-5716-428d-9cf5-b4a38a6f3c9c",
            "Name": "Other",
            "Username": "other",
            "Email": "other@mail.com",
            "Password": "$2a$04$Kx1Ir6X8S21m5DxxnqD5BeESSv01PZ3E5jnvRA954WH3GpCtmiWby",
            "SocialId": "",
            "Provider": "",
            "Avatar": "",
            "RoleId": "f7dfeabb-3a5f-415f-8b84-785b87767760",
            "CreatedAt": "2022-08-25T09:09:50.54+07:00",
            "UpdatedAt": "2022-08-25T09:09:50.54+07:00",
            "DeletedAt": null
        },
        "message": "Data created successfully",
        "status": true
    }
    ```

3. Update User
    #### # Postman
    #### Method: PUT
    #### Authorization: Bearer Token *(token)*
    #### Url: <http://localhost:8000/api/v1/user/{userId}>
    #### Body, raw, json
    ```json
    {
        "name": "Other",
        "username": "other",
        "email": "other@mail.com",
        "role_id": "f7dfeabb-3a5f-415f-8b84-785b87767760"
    }
    ```

    #### Response
    ```json
    {
        "data": {
            "Id": "73907481-5716-428d-9cf5-b4a38a6f3c9c",
            "Name": "Other",
            "Username": "other",
            "Email": "other@mail.com",
            "Password": "$2a$04$Kx1Ir6X8S21m5DxxnqD5BeESSv01PZ3E5jnvRA954WH3GpCtmiWby",
            "SocialId": "",
            "Provider": "",
            "Avatar": "",
            "RoleId": "f7dfeabb-3a5f-415f-8b84-785b87767760",
            "CreatedAt": "2022-08-25T09:09:50.54+07:00",
            "UpdatedAt": "2022-08-25T09:09:50.54+07:00",
            "DeletedAt": null
        },
        "message": "Data created successfully",
        "status": true
    }
    ```

4. Delete User
    #### # Postman
    #### Method: DELETE
    #### Authorization: Bearer Token *(token)*
    #### Url: <http://localhost:8000/api/v1/user/{UserId}>
    #### Response
    ```json
    {
        "data": null,
        "message": "Data deleted successfully",
        "status": true
    }
    ```

## # JWT

1. Refresh Token
    #### # Postman
    #### Method: POST
    #### Authorization: Bearer Token *(token)*
    #### Url: <http://localhost:8000/api/v1/auth/refresh-token>
    #### Body, raw, json
    ```json
    {
        "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjE1NjUxOTEsImlhdCI6MTY2MTM5MjM5MSwicmVmcmVzaF9pZCI6IjI5NzJjZTRlLTY4YTAtNDI5Ny1hZjZhLTFhMmI5M2Y1NTlhYSJ9.KyFRWbG8MbsX6qGmzdWRQqEwTYdOddMkNmcA1iL2alI"
    }
    ```

    #### Response
    ```json
    {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjE0ODAzMDksImlhdCI6MTY2MTM5MzkwOSwidXNlcl9pZCI6ImQ4Mzg4MDI3LTQ5ZmYtNDAzYy1iYWQyLTIzNGY2MzI1NmE4NyJ9.jXclCvOuEcLEyS3RTImoAINbrtoudWLCAL3dOCdDEWk"
    }
    ```

2. Revoke Refresh Token
    #### # Postman
    #### Method: PUT
    #### Authorization: Bearer Token *(token)*
    #### Url: <http://localhost:8000/api/v1/revoke>
    #### Body, raw, json
    ```json
    {
        "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjE1NjUxOTEsImlhdCI6MTY2MTM5MjM5MSwicmVmcmVzaF9pZCI6IjI5NzJjZTRlLTY4YTAtNDI5Ny1hZjZhLTFhMmI5M2Y1NTlhYSJ9.KyFRWbG8MbsX6qGmzdWRQqEwTYdOddMkNmcA1iL2alI"
    }
    ```

    #### Response
    ```json
    {
        "data": null,
        "message": "Revoke refresh token successfully",
        "status": true
    }
    ```

## # File Management

1. Upload File
    #### # Postman
    #### Method: POST
    #### Authorization: Bearer Token *(token)*
    #### Url: <http://localhost:8000/api/v1/upload/public>
    #### *disk= public/private*
    #### Url: <http://localhost:8000/api/v1/upload/{public/private}>
    #### Body, form-data
    ```
    KEY     :   VALUE
    file    :   5.jpg
    ```

    #### Response
    ```json
    {
        "data": "4.jpg",
        "message": "Success",
        "status": true
    }
    ```

2. Show File Public Encode
    #### # Postman
    #### Method: GET
    #### Authorization: Bearer Token *(token)*
    #### Url: <http://localhost:8000/api/v1/encode/4.jpg>
    #### Response
    ```json
    {
        "data": "data:image/jpeg;base64,....",
        "message": null,
        "status": true
    }
    ```
3. Show File Stream Public
    #### # Postman
    #### Method: GET
    #### Authorization: Bearer Token *(token)*
    #### Url: <http://localhost:8000/api/v1/storage/4.jpg>
    #### Response
    ```json
    File Image
    ```

## # Change Password

#### # Postman
#### Method: PUT
#### Authorization: Bearer Token *(token)*
#### Url: <http://localhost:8000/api/v1/changepassword>
#### Body, raw, json
```json
{
    "old_password": "12345678",
    "new_password": "1234567890"
}
```

#### Response
```json
{
    "data": {
        "Id": "d8388027-49ff-403c-bad2-234f63256a87",
        "Name": "Super Admin",
        "Username": "superadmin",
        "Email": "superadmin@mail.com",
        "Password": "$2a$04$Ee5gPdT9/ue6POxotbFcFOt3mRiZ7H9GnKH22.lHc7LdsJ27lp6mK",
        "SocialId": "",
        "Provider": "",
        "Avatar": "",
        "RoleId": "f7dfeabb-3a5f-415f-8b84-785b87767760",
        "CreatedAt": "2022-08-24T16:00:28.685+07:00",
        "UpdatedAt": "2022-08-25T09:55:38.398+07:00",
        "DeletedAt": null
    },
    "message": "Change password successfully",
    "status": true
}
```

## # Socialite

1. Get url with browser

    ```json
    http://localhost:8000/api/v1/auth/google
    http://localhost:8000/api/v1/auth/github
    ```

2. Select account & login
3. Response

    ```json
    {
        "data": {
            "Id": "458549c7-ae2c-42b7-b0c1-7011309dd28c",
            "Name": "User Name",
            "Username": "Username",
            "Email": "user@gmail.com",
            "Password": "",
            "SocialId": "11810....",
            "Provider": "google",
            "Avatar": "https://lh3.googleusercontent.com/a-/...",
            "RoleId": "",
            "CreatedAt": "2022-08-25T10:13:34.745+07:00",
            "UpdatedAt": "2022-08-25T10:13:34.745+07:00",
            "DeletedAt": null
        },
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjE0ODM2MTQsImlhdCI6MTY2MTM5NzIxNCwidXNlcl9pZCI6IjQ1ODU0OWM3LWFlMmMtNDJiNy1iMGMxLTcwMTEzMDlkZDI4YyJ9...."
    }
    ```
