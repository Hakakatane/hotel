package hotel

import (
	"fmt"
	"example.com/structure"
)

type Company interface {
	Slogan()
}

// Отель
type Hotel struct {
	Structure     structure.Structure
	Name          string
	Stars         int
	PricePerNight float64
	Availability  string // места есть, мест нет, забронирован, работает, закрыто
}

// Конструктор для Hotel
func NewHotel(s structure.Structure, name string, stars int, price float64) Hotel {
	return Hotel{
		Structure:     s,
		Name:          name,
		Stars:         stars,
		PricePerNight: price,
		Availability:  "места есть",
	}
}

// Методы для изменения состояния
func (h *Hotel) Available() {
	h.Availability = "места есть"
	fmt.Println("В отеле есть свободные места.")
}

func (h *Hotel) NoRooms() {
	h.Availability = "мест нет"
	fmt.Println("В отеле нет свободных мест.")
}

func (h *Hotel) Booked() {
	h.Availability = "забронирован"
	fmt.Println("Отель забронирован.")
}

func (h *Hotel) Open() {
	h.Availability = "работает"
	fmt.Println("Отель работает.")
}

func (h *Hotel) Close() {
	h.Availability = "закрыто"
	fmt.Println("Отель закрыт.")
}

func (h *Hotel) SetAvailability(availability string) {
	h.Availability = availability
	fmt.Printf("Доступность отеля изменена на %s\n", availability)
}

// Методы для получения характеристик
func (h Hotel) GetName() string {
	return h.Name
}

func (h Hotel) GetStars() int {
	return h.Stars
}

func (h Hotel) GetPricePerNight() float64 {
	return h.PricePerNight
}

func (h Hotel) GetAvailability() string {
	return h.Availability
}

func (h *Hotel) Slogan() {
	fmt.Printf("Больше - лучше\n")
}
