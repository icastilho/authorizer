// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type AddEmailTemplateRequest struct {
	EventName string  `json:"event_name"`
	Subject   string  `json:"subject"`
	Template  string  `json:"template"`
	Design    *string `json:"design"`
}

type AddWebhookRequest struct {
	EventName        string                 `json:"event_name"`
	EventDescription *string                `json:"event_description"`
	Endpoint         string                 `json:"endpoint"`
	Enabled          bool                   `json:"enabled"`
	Headers          map[string]interface{} `json:"headers"`
}

type AdminLoginInput struct {
	AdminSecret string `json:"admin_secret"`
}

type AdminSignupInput struct {
	AdminSecret string `json:"admin_secret"`
}

type AuthResponse struct {
	Message                   string  `json:"message"`
	ShouldShowEmailOtpScreen  *bool   `json:"should_show_email_otp_screen"`
	ShouldShowMobileOtpScreen *bool   `json:"should_show_mobile_otp_screen"`
	AccessToken               *string `json:"access_token"`
	IDToken                   *string `json:"id_token"`
	RefreshToken              *string `json:"refresh_token"`
	ExpiresIn                 *int64  `json:"expires_in"`
	User                      *User   `json:"user"`
}

type DeleteEmailTemplateRequest struct {
	ID string `json:"id"`
}

type DeleteUserInput struct {
	Email string `json:"email"`
}

type EmailTemplate struct {
	ID        string `json:"id"`
	EventName string `json:"event_name"`
	Template  string `json:"template"`
	Design    string `json:"design"`
	Subject   string `json:"subject"`
	CreatedAt *int64 `json:"created_at"`
	UpdatedAt *int64 `json:"updated_at"`
}

type EmailTemplates struct {
	Pagination     *Pagination      `json:"pagination"`
	EmailTemplates []*EmailTemplate `json:"email_templates"`
}

type Env struct {
	AccessTokenExpiryTime            *string  `json:"ACCESS_TOKEN_EXPIRY_TIME"`
	AdminSecret                      *string  `json:"ADMIN_SECRET"`
	DatabaseName                     *string  `json:"DATABASE_NAME"`
	DatabaseURL                      *string  `json:"DATABASE_URL"`
	DatabaseType                     *string  `json:"DATABASE_TYPE"`
	DatabaseUsername                 *string  `json:"DATABASE_USERNAME"`
	DatabasePassword                 *string  `json:"DATABASE_PASSWORD"`
	DatabaseHost                     *string  `json:"DATABASE_HOST"`
	DatabasePort                     *string  `json:"DATABASE_PORT"`
	ClientID                         string   `json:"CLIENT_ID"`
	ClientSecret                     string   `json:"CLIENT_SECRET"`
	CustomAccessTokenScript          *string  `json:"CUSTOM_ACCESS_TOKEN_SCRIPT"`
	SMTPHost                         *string  `json:"SMTP_HOST"`
	SMTPPort                         *string  `json:"SMTP_PORT"`
	SMTPUsername                     *string  `json:"SMTP_USERNAME"`
	SMTPPassword                     *string  `json:"SMTP_PASSWORD"`
	SMTPLocalName                    *string  `json:"SMTP_LOCAL_NAME"`
	SenderEmail                      *string  `json:"SENDER_EMAIL"`
	SenderName                       *string  `json:"SENDER_NAME"`
	JwtType                          *string  `json:"JWT_TYPE"`
	JwtSecret                        *string  `json:"JWT_SECRET"`
	JwtPrivateKey                    *string  `json:"JWT_PRIVATE_KEY"`
	JwtPublicKey                     *string  `json:"JWT_PUBLIC_KEY"`
	AllowedOrigins                   []string `json:"ALLOWED_ORIGINS"`
	AppURL                           *string  `json:"APP_URL"`
	RedisURL                         *string  `json:"REDIS_URL"`
	ResetPasswordURL                 *string  `json:"RESET_PASSWORD_URL"`
	DisableEmailVerification         bool     `json:"DISABLE_EMAIL_VERIFICATION"`
	DisableBasicAuthentication       bool     `json:"DISABLE_BASIC_AUTHENTICATION"`
	DisableMagicLinkLogin            bool     `json:"DISABLE_MAGIC_LINK_LOGIN"`
	DisableLoginPage                 bool     `json:"DISABLE_LOGIN_PAGE"`
	DisableSignUp                    bool     `json:"DISABLE_SIGN_UP"`
	DisableRedisForEnv               bool     `json:"DISABLE_REDIS_FOR_ENV"`
	DisableStrongPassword            bool     `json:"DISABLE_STRONG_PASSWORD"`
	DisableMultiFactorAuthentication bool     `json:"DISABLE_MULTI_FACTOR_AUTHENTICATION"`
	EnforceMultiFactorAuthentication bool     `json:"ENFORCE_MULTI_FACTOR_AUTHENTICATION"`
	Roles                            []string `json:"ROLES"`
	ProtectedRoles                   []string `json:"PROTECTED_ROLES"`
	DefaultRoles                     []string `json:"DEFAULT_ROLES"`
	JwtRoleClaim                     *string  `json:"JWT_ROLE_CLAIM"`
	GoogleClientID                   *string  `json:"GOOGLE_CLIENT_ID"`
	GoogleClientSecret               *string  `json:"GOOGLE_CLIENT_SECRET"`
	GithubClientID                   *string  `json:"GITHUB_CLIENT_ID"`
	GithubClientSecret               *string  `json:"GITHUB_CLIENT_SECRET"`
	FacebookClientID                 *string  `json:"FACEBOOK_CLIENT_ID"`
	FacebookClientSecret             *string  `json:"FACEBOOK_CLIENT_SECRET"`
	LinkedinClientID                 *string  `json:"LINKEDIN_CLIENT_ID"`
	LinkedinClientSecret             *string  `json:"LINKEDIN_CLIENT_SECRET"`
	AppleClientID                    *string  `json:"APPLE_CLIENT_ID"`
	AppleClientSecret                *string  `json:"APPLE_CLIENT_SECRET"`
	TwitterClientID                  *string  `json:"TWITTER_CLIENT_ID"`
	TwitterClientSecret              *string  `json:"TWITTER_CLIENT_SECRET"`
	MicrosoftClientID                *string  `json:"MICROSOFT_CLIENT_ID"`
	MicrosoftClientSecret            *string  `json:"MICROSOFT_CLIENT_SECRET"`
	MicrosoftActiveDirectoryTenantID *string  `json:"MICROSOFT_ACTIVE_DIRECTORY_TENANT_ID"`
	OrganizationName                 *string  `json:"ORGANIZATION_NAME"`
	OrganizationLogo                 *string  `json:"ORGANIZATION_LOGO"`
	AppCookieSecure                  bool     `json:"APP_COOKIE_SECURE"`
	AdminCookieSecure                bool     `json:"ADMIN_COOKIE_SECURE"`
	DefaultAuthorizeResponseType     *string  `json:"DEFAULT_AUTHORIZE_RESPONSE_TYPE"`
	DefaultAuthorizeResponseMode     *string  `json:"DEFAULT_AUTHORIZE_RESPONSE_MODE"`
}

type Error struct {
	Message string `json:"message"`
	Reason  string `json:"reason"`
}

type ForgotPasswordInput struct {
	Email       string  `json:"email"`
	State       *string `json:"state"`
	RedirectURI *string `json:"redirect_uri"`
}

type GenerateJWTKeysInput struct {
	Type string `json:"type"`
}

type GenerateJWTKeysResponse struct {
	Secret     *string `json:"secret"`
	PublicKey  *string `json:"public_key"`
	PrivateKey *string `json:"private_key"`
}

type GetUserRequest struct {
	ID    *string `json:"id"`
	Email *string `json:"email"`
}

type InviteMemberInput struct {
	Emails      []string `json:"emails"`
	RedirectURI *string  `json:"redirect_uri"`
}

type InviteMembersResponse struct {
	Message string  `json:"message"`
	Users   []*User `json:"Users"`
}

type ListWebhookLogRequest struct {
	Pagination *PaginationInput `json:"pagination"`
	WebhookID  *string          `json:"webhook_id"`
}

type LoginInput struct {
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Roles    []string `json:"roles"`
	Scope    []string `json:"scope"`
	State    *string  `json:"state"`
}

type MagicLinkLoginInput struct {
	Email       string   `json:"email"`
	Roles       []string `json:"roles"`
	Scope       []string `json:"scope"`
	State       *string  `json:"state"`
	RedirectURI *string  `json:"redirect_uri"`
}

type Meta struct {
	Version                      string `json:"version"`
	ClientID                     string `json:"client_id"`
	IsGoogleLoginEnabled         bool   `json:"is_google_login_enabled"`
	IsFacebookLoginEnabled       bool   `json:"is_facebook_login_enabled"`
	IsGithubLoginEnabled         bool   `json:"is_github_login_enabled"`
	IsLinkedinLoginEnabled       bool   `json:"is_linkedin_login_enabled"`
	IsAppleLoginEnabled          bool   `json:"is_apple_login_enabled"`
	IsTwitterLoginEnabled        bool   `json:"is_twitter_login_enabled"`
	IsMicrosoftLoginEnabled      bool   `json:"is_microsoft_login_enabled"`
	IsEmailVerificationEnabled   bool   `json:"is_email_verification_enabled"`
	IsBasicAuthenticationEnabled bool   `json:"is_basic_authentication_enabled"`
	IsMagicLinkLoginEnabled      bool   `json:"is_magic_link_login_enabled"`
	IsSignUpEnabled              bool   `json:"is_sign_up_enabled"`
	IsStrongPasswordEnabled      bool   `json:"is_strong_password_enabled"`
	IsMultiFactorAuthEnabled     bool   `json:"is_multi_factor_auth_enabled"`
}

type MobileLoginInput struct {
	PhoneNumber string   `json:"phone_number"`
	Password    string   `json:"password"`
	Roles       []string `json:"roles"`
	Scope       []string `json:"scope"`
	State       *string  `json:"state"`
}

type MobileSignUpInput struct {
	Email                    *string  `json:"email"`
	GivenName                *string  `json:"given_name"`
	FamilyName               *string  `json:"family_name"`
	MiddleName               *string  `json:"middle_name"`
	Nickname                 *string  `json:"nickname"`
	Gender                   *string  `json:"gender"`
	Birthdate                *string  `json:"birthdate"`
	PhoneNumber              string   `json:"phone_number"`
	Picture                  *string  `json:"picture"`
	Password                 string   `json:"password"`
	ConfirmPassword          string   `json:"confirm_password"`
	Roles                    []string `json:"roles"`
	Scope                    []string `json:"scope"`
	RedirectURI              *string  `json:"redirect_uri"`
	IsMultiFactorAuthEnabled *bool    `json:"is_multi_factor_auth_enabled"`
	State                    *string  `json:"state"`
}

type OAuthRevokeInput struct {
	RefreshToken string `json:"refresh_token"`
}

type PaginatedInput struct {
	Pagination *PaginationInput `json:"pagination"`
}

type Pagination struct {
	Limit  int64 `json:"limit"`
	Page   int64 `json:"page"`
	Offset int64 `json:"offset"`
	Total  int64 `json:"total"`
}

type PaginationInput struct {
	Limit *int64 `json:"limit"`
	Page  *int64 `json:"page"`
}

type ResendOTPRequest struct {
	Email       *string `json:"email"`
	PhoneNumber *string `json:"phone_number"`
	State       *string `json:"state"`
}

type ResendVerifyEmailInput struct {
	Email      string  `json:"email"`
	Identifier string  `json:"identifier"`
	State      *string `json:"state"`
}

type ResetPasswordInput struct {
	Token           string `json:"token"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type Response struct {
	Message string `json:"message"`
}

type SMSVerificationRequests struct {
	ID            string `json:"id"`
	Code          string `json:"code"`
	CodeExpiresAt int64  `json:"code_expires_at"`
	PhoneNumber   string `json:"phone_number"`
	CreatedAt     int64  `json:"created_at"`
	UpdatedAt     *int64 `json:"updated_at"`
}

type SessionQueryInput struct {
	Roles []string `json:"roles"`
	Scope []string `json:"scope"`
}

type SignUpInput struct {
	Email                    string   `json:"email"`
	GivenName                *string  `json:"given_name"`
	FamilyName               *string  `json:"family_name"`
	MiddleName               *string  `json:"middle_name"`
	Nickname                 *string  `json:"nickname"`
	Gender                   *string  `json:"gender"`
	Birthdate                *string  `json:"birthdate"`
	PhoneNumber              *string  `json:"phone_number"`
	Picture                  *string  `json:"picture"`
	Password                 string   `json:"password"`
	ConfirmPassword          string   `json:"confirm_password"`
	Roles                    []string `json:"roles"`
	Scope                    []string `json:"scope"`
	RedirectURI              *string  `json:"redirect_uri"`
	IsMultiFactorAuthEnabled *bool    `json:"is_multi_factor_auth_enabled"`
	State                    *string  `json:"state"`
}

type TestEndpointRequest struct {
	Endpoint         string                 `json:"endpoint"`
	EventName        string                 `json:"event_name"`
	EventDescription *string                `json:"event_description"`
	Headers          map[string]interface{} `json:"headers"`
}

type TestEndpointResponse struct {
	HTTPStatus *int64  `json:"http_status"`
	Response   *string `json:"response"`
}

type UpdateAccessInput struct {
	UserID string `json:"user_id"`
}

type UpdateEmailTemplateRequest struct {
	ID        string  `json:"id"`
	EventName *string `json:"event_name"`
	Template  *string `json:"template"`
	Subject   *string `json:"subject"`
	Design    *string `json:"design"`
}

type UpdateEnvInput struct {
	AccessTokenExpiryTime            *string  `json:"ACCESS_TOKEN_EXPIRY_TIME"`
	AdminSecret                      *string  `json:"ADMIN_SECRET"`
	CustomAccessTokenScript          *string  `json:"CUSTOM_ACCESS_TOKEN_SCRIPT"`
	OldAdminSecret                   *string  `json:"OLD_ADMIN_SECRET"`
	SMTPHost                         *string  `json:"SMTP_HOST"`
	SMTPPort                         *string  `json:"SMTP_PORT"`
	SMTPUsername                     *string  `json:"SMTP_USERNAME"`
	SMTPPassword                     *string  `json:"SMTP_PASSWORD"`
	SMTPLocalName                    *string  `json:"SMTP_LOCAL_NAME"`
	SenderEmail                      *string  `json:"SENDER_EMAIL"`
	SenderName                       *string  `json:"SENDER_NAME"`
	JwtType                          *string  `json:"JWT_TYPE"`
	JwtSecret                        *string  `json:"JWT_SECRET"`
	JwtPrivateKey                    *string  `json:"JWT_PRIVATE_KEY"`
	JwtPublicKey                     *string  `json:"JWT_PUBLIC_KEY"`
	AllowedOrigins                   []string `json:"ALLOWED_ORIGINS"`
	AppURL                           *string  `json:"APP_URL"`
	ResetPasswordURL                 *string  `json:"RESET_PASSWORD_URL"`
	AppCookieSecure                  *bool    `json:"APP_COOKIE_SECURE"`
	AdminCookieSecure                *bool    `json:"ADMIN_COOKIE_SECURE"`
	DisableEmailVerification         *bool    `json:"DISABLE_EMAIL_VERIFICATION"`
	DisableBasicAuthentication       *bool    `json:"DISABLE_BASIC_AUTHENTICATION"`
	DisableMagicLinkLogin            *bool    `json:"DISABLE_MAGIC_LINK_LOGIN"`
	DisableLoginPage                 *bool    `json:"DISABLE_LOGIN_PAGE"`
	DisableSignUp                    *bool    `json:"DISABLE_SIGN_UP"`
	DisableRedisForEnv               *bool    `json:"DISABLE_REDIS_FOR_ENV"`
	DisableStrongPassword            *bool    `json:"DISABLE_STRONG_PASSWORD"`
	DisableMultiFactorAuthentication *bool    `json:"DISABLE_MULTI_FACTOR_AUTHENTICATION"`
	EnforceMultiFactorAuthentication *bool    `json:"ENFORCE_MULTI_FACTOR_AUTHENTICATION"`
	Roles                            []string `json:"ROLES"`
	ProtectedRoles                   []string `json:"PROTECTED_ROLES"`
	DefaultRoles                     []string `json:"DEFAULT_ROLES"`
	JwtRoleClaim                     *string  `json:"JWT_ROLE_CLAIM"`
	GoogleClientID                   *string  `json:"GOOGLE_CLIENT_ID"`
	GoogleClientSecret               *string  `json:"GOOGLE_CLIENT_SECRET"`
	GithubClientID                   *string  `json:"GITHUB_CLIENT_ID"`
	GithubClientSecret               *string  `json:"GITHUB_CLIENT_SECRET"`
	FacebookClientID                 *string  `json:"FACEBOOK_CLIENT_ID"`
	FacebookClientSecret             *string  `json:"FACEBOOK_CLIENT_SECRET"`
	LinkedinClientID                 *string  `json:"LINKEDIN_CLIENT_ID"`
	LinkedinClientSecret             *string  `json:"LINKEDIN_CLIENT_SECRET"`
	AppleClientID                    *string  `json:"APPLE_CLIENT_ID"`
	AppleClientSecret                *string  `json:"APPLE_CLIENT_SECRET"`
	TwitterClientID                  *string  `json:"TWITTER_CLIENT_ID"`
	TwitterClientSecret              *string  `json:"TWITTER_CLIENT_SECRET"`
	MicrosoftClientID                *string  `json:"MICROSOFT_CLIENT_ID"`
	MicrosoftClientSecret            *string  `json:"MICROSOFT_CLIENT_SECRET"`
	MicrosoftActiveDirectoryTenantID *string  `json:"MICROSOFT_ACTIVE_DIRECTORY_TENANT_ID"`
	OrganizationName                 *string  `json:"ORGANIZATION_NAME"`
	OrganizationLogo                 *string  `json:"ORGANIZATION_LOGO"`
	DefaultAuthorizeResponseType     *string  `json:"DEFAULT_AUTHORIZE_RESPONSE_TYPE"`
	DefaultAuthorizeResponseMode     *string  `json:"DEFAULT_AUTHORIZE_RESPONSE_MODE"`
}

type UpdateProfileInput struct {
	OldPassword              *string `json:"old_password"`
	NewPassword              *string `json:"new_password"`
	ConfirmNewPassword       *string `json:"confirm_new_password"`
	Email                    *string `json:"email"`
	GivenName                *string `json:"given_name"`
	FamilyName               *string `json:"family_name"`
	MiddleName               *string `json:"middle_name"`
	Nickname                 *string `json:"nickname"`
	Gender                   *string `json:"gender"`
	Birthdate                *string `json:"birthdate"`
	PhoneNumber              *string `json:"phone_number"`
	Picture                  *string `json:"picture"`
	IsMultiFactorAuthEnabled *bool   `json:"is_multi_factor_auth_enabled"`
}

type UpdateUserInput struct {
	ID                       string    `json:"id"`
	Email                    *string   `json:"email"`
	EmailVerified            *bool     `json:"email_verified"`
	GivenName                *string   `json:"given_name"`
	FamilyName               *string   `json:"family_name"`
	MiddleName               *string   `json:"middle_name"`
	Nickname                 *string   `json:"nickname"`
	Gender                   *string   `json:"gender"`
	Birthdate                *string   `json:"birthdate"`
	PhoneNumber              *string   `json:"phone_number"`
	Picture                  *string   `json:"picture"`
	Roles                    []*string `json:"roles"`
	IsMultiFactorAuthEnabled *bool     `json:"is_multi_factor_auth_enabled"`
}

type UpdateWebhookRequest struct {
	ID               string                 `json:"id"`
	EventName        *string                `json:"event_name"`
	EventDescription *string                `json:"event_description"`
	Endpoint         *string                `json:"endpoint"`
	Enabled          *bool                  `json:"enabled"`
	Headers          map[string]interface{} `json:"headers"`
}

type User struct {
	ID                       string   `json:"id"`
	Email                    string   `json:"email"`
	EmailVerified            bool     `json:"email_verified"`
	SignupMethods            string   `json:"signup_methods"`
	GivenName                *string  `json:"given_name"`
	FamilyName               *string  `json:"family_name"`
	MiddleName               *string  `json:"middle_name"`
	Nickname                 *string  `json:"nickname"`
	PreferredUsername        *string  `json:"preferred_username"`
	Gender                   *string  `json:"gender"`
	Birthdate                *string  `json:"birthdate"`
	PhoneNumber              *string  `json:"phone_number"`
	PhoneNumberVerified      *bool    `json:"phone_number_verified"`
	Picture                  *string  `json:"picture"`
	Roles                    []string `json:"roles"`
	CreatedAt                *int64   `json:"created_at"`
	UpdatedAt                *int64   `json:"updated_at"`
	RevokedTimestamp         *int64   `json:"revoked_timestamp"`
	IsMultiFactorAuthEnabled *bool    `json:"is_multi_factor_auth_enabled"`
}

type Users struct {
	Pagination *Pagination `json:"pagination"`
	Users      []*User     `json:"users"`
}

type ValidateJWTTokenInput struct {
	TokenType string   `json:"token_type"`
	Token     string   `json:"token"`
	Roles     []string `json:"roles"`
}

type ValidateJWTTokenResponse struct {
	IsValid bool                   `json:"is_valid"`
	Claims  map[string]interface{} `json:"claims"`
}

type ValidateSessionInput struct {
	Cookie string   `json:"cookie"`
	Roles  []string `json:"roles"`
}

type ValidateSessionResponse struct {
	IsValid bool `json:"is_valid"`
}

type VerificationRequest struct {
	ID          string  `json:"id"`
	Identifier  *string `json:"identifier"`
	Token       *string `json:"token"`
	Email       *string `json:"email"`
	Expires     *int64  `json:"expires"`
	CreatedAt   *int64  `json:"created_at"`
	UpdatedAt   *int64  `json:"updated_at"`
	Nonce       *string `json:"nonce"`
	RedirectURI *string `json:"redirect_uri"`
}

type VerificationRequests struct {
	Pagination           *Pagination            `json:"pagination"`
	VerificationRequests []*VerificationRequest `json:"verification_requests"`
}

type VerifyEmailInput struct {
	Token string  `json:"token"`
	State *string `json:"state"`
}

type VerifyOTPRequest struct {
	Email       *string `json:"email"`
	PhoneNumber *string `json:"phone_number"`
	Otp         string  `json:"otp"`
	State       *string `json:"state"`
}

type Webhook struct {
	ID               string                 `json:"id"`
	EventName        *string                `json:"event_name"`
	EventDescription *string                `json:"event_description"`
	Endpoint         *string                `json:"endpoint"`
	Enabled          *bool                  `json:"enabled"`
	Headers          map[string]interface{} `json:"headers"`
	CreatedAt        *int64                 `json:"created_at"`
	UpdatedAt        *int64                 `json:"updated_at"`
}

type WebhookLog struct {
	ID         string  `json:"id"`
	HTTPStatus *int64  `json:"http_status"`
	Response   *string `json:"response"`
	Request    *string `json:"request"`
	WebhookID  *string `json:"webhook_id"`
	CreatedAt  *int64  `json:"created_at"`
	UpdatedAt  *int64  `json:"updated_at"`
}

type WebhookLogs struct {
	Pagination  *Pagination   `json:"pagination"`
	WebhookLogs []*WebhookLog `json:"webhook_logs"`
}

type WebhookRequest struct {
	ID string `json:"id"`
}

type Webhooks struct {
	Pagination *Pagination `json:"pagination"`
	Webhooks   []*Webhook  `json:"webhooks"`
}
