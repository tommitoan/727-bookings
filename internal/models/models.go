package models

import (
	"time"
)

// User is the user model
type User struct {
	ID            int
	UserName      string
	Email         string
	Password      string
	PaymentID     int
	CustomerID    int
	ReservationID int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// UserPayment is the User payment model
type UserPayment struct {
	ID          int
	PaymentType string
	Amount      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Customer is the customer model
type Customer struct {
	ID            int
	UserID        int
	FirstName     string
	LastName      string
	Gender        string
	Birth         time.Time
	SolarCalendar bool
	Intercalation bool
	FatherZodiac  string
	MotherZodiac  string
	AstrologyID   int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// Horoscope is the Customer horoscope model
type Horoscope struct {
	ID          int
	CustomerID  int
	YearPillar  string
	MonthPillar string
	DayPillar   string
	HourPillar  string
	AstrologyID int
	SummaryID   int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// HoroscopeSummary is the Horoscope summary model
type HoroscopeSummary struct {
	ID          int
	HoroscopeID int
	ServiceID   int
	Content     string
	Paid        bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Reservation is the reservation model
type Reservation struct {
	ID            int
	UserID        int
	CustomerID    int
	RestrictionID int
	ServiceID     int
	BookingDate   time.Time
	BookingType   string
	BookingPlace  string
	Cost          int
	Deposit       int
	Status        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// Restriction is the reservation restriction model
type Restriction struct {
	ID              int
	RestrictionName string
	BookingTime     time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// Service is the Service model
type Service struct {
	ID          int
	ServiceName string
	Description string
	Duration    time.Duration
	Price       int
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// MailData holds an email message
type MailData struct {
	To       string
	From     string
	Subject  string
	Content  string
	Template string
}
