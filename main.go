package main

import (
	"fmt"
	"sort"
)

// Event struct to represent an event with start and end time.
type Event struct {
	StartTime int
	EndTime   int
}

// Scheduler struct to manage a list of events.
type Scheduler struct {
	events []Event
}

// NewScheduler creates and returns a new Scheduler instance.
func NewScheduler() *Scheduler {
	return &Scheduler{events: []Event{}}
}

// AddEvent adds a new event if it doesn't overlap with existing events.
// Returns true if the event is added successfully, false if it overlaps.
func (s *Scheduler) AddEvent(newEvent Event) bool {
	for _, event := range s.events {
		if isOverlap(event, newEvent) {
			return false
		}
	}
	s.events = append(s.events, newEvent)
	sort.Slice(s.events, func(i, j int) bool {
		return s.events[i].StartTime < s.events[j].StartTime
	})
	return true
}

// GetEvents returns all scheduled events.
func (s *Scheduler) GetEvents() []Event {
	return s.events
}

// Helper function to check if two events overlap.
func isOverlap(e1, e2 Event) bool {
	return !(e1.EndTime <= e2.StartTime || e2.EndTime <= e1.StartTime)
}

// DisplayEvents prints all events in a readable format.
func (s *Scheduler) DisplayEvents() {
	fmt.Println("Scheduled Events:")
	for _, event := range s.events {
		fmt.Printf("Event: %02d:00 - %02d:00\n", event.StartTime, event.EndTime)
	}
}

func main() {
	scheduler := NewScheduler()

	for {
		fmt.Println("\n1. Add Event")
		fmt.Println("2. View Events")
		fmt.Println("3. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var start, end int
			fmt.Print("Enter start time (0-23): ")
			fmt.Scan(&start)
			fmt.Print("Enter end time (1-24): ")
			fmt.Scan(&end)

			if start >= end || start < 0 || end > 24 {
				fmt.Println("Invalid time range! Try again.")
				continue
			}

			if scheduler.AddEvent(Event{StartTime: start, EndTime: end}) {
				fmt.Println("Event added successfully!")
			} else {
				fmt.Println("Failed to add event! Overlapping with an existing event.")
			}

		case 2:
			scheduler.DisplayEvents()

		case 3:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice! Try again.")
		}
	}
}
