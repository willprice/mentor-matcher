package domain

type UserProfileRepository interface {
	CreateUserProfile(profile *UserProfile) (int, error)
	GetUserProfileByEmail(email string) (*UserProfile, error)
}
