package challenge

import (
	"log"
	"math"
	"time"
)

type Subscription struct {
	Id                    int
	CustomerId            int
	MonthlyPriceInDollars int
}

type User struct {
	Id            int
	Name          string
	ActivatedOn   time.Time
	DeactivatedOn time.Time
	CustomerId    int
}

// Computes the monthly charge for a given subscription.
//
// Returns the total monthly bill for the customer in dollars and cents, rounded
// to 2 decimal places.
// If there are no active users or the subscription is null, returns 0.
//
// month: Always present
//   Has the following structure:
//   "2022-04"  // April 2022 in YYYY-MM format
//
// activeSubscription: May be null
//   If present, has the following structure (see Subscription struct):
//   {
//     Id: 763,
//     CustomerId: 328,
//     MonthlyPriceInDollars: 4  // price per active user per month
//   }
//
// users: May be empty, but not null
//   Has the following structure (see User struct):
//   {
//     {
//       Id: 1,
//       Name: "Employee #1",
//       CustomerId: 1,
//
//       // when this user started
//       Activated_On: time.Date(2021, 11, 4, 0, 0, 0, 0, time.UTC),
//
//       // last day to bill for user
//       // should bill up to and including this date
//       // since user had some access on this date
//       DeactivatedOn: time.Date(2022, 4, 10, 0, 0, 0, 0, time.UTC)
//     },
//     {
//       Id: 2,
//       Name: "Employee #2",
//       CustomerId: 1,
//
//       // when this user started
//       ActivatedOn: time.Date(2021, 12, 4, 0, 0, 0, 0, time.UTC),
//
//       // hasn't been deactivated yet
//       DeactivatedOn: nil
//     },
//   }
func BillFor(yearMonth string, activeSubscription *Subscription, users *[]User) float64 {
	// TODO: your code here

	month, err := time.Parse("%m", yearMonth)
	if err != nil {
		log.Fatal(err)
	}

	firstDay := FirstDayOfMonth(month)
	lastDay := LastDayOfMonth(month)

	daysInMonth := lastDay.Sub(firstDay).Hours()

	if activeSubscription == nil {
		return 0
	}

	var dailyRate float64 = 0
	cost := activeSubscription.MonthlyPriceInDollars
	dailyRate += float64(cost) / (daysInMonth)
	var totalCost float64 = 0

	for _, user := range *users {
		dDate := user.DeactivatedOn
		aDate := user.ActivatedOn

		// 		 if dDate is not None:
		if dDate.IsDST() {
			if dDate.Before(month) {
				totalCost += 0
			} else {
				activeDays := float64(dDate.Sub(firstDay)) + 1
				totalCost += activeDays * dailyRate
			}
		} else

		//        if dDate is None:
		if aDate.After(firstDay) || aDate.Equal(firstDay) && aDate.Before(lastDay) {
			activeDays := int(aDate.Sub(lastDay)) + 1
			totalCost += float64(activeDays) * (dailyRate)
		} else {
			totalCost += daysInMonth * dailyRate
		}
	}

	return math.Round(totalCost)
}

/*******************
* Helper functions *
*******************/

/*
Takes a time.Time object and returns a time.Time object
which is the first day of that month.

FirstDayOfMonth(time.Date(2019, 2, 7, 0, 0, 0, 0, time.UTC))  // Feb 7
=> time.Date(2019, 2, 1, 0, 0, 0, 0, time.UTC))               // Feb 1
*/
func FirstDayOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

/*
Takes a time.Time object and returns a time.Time object
which is the end of the last day of that month.

LastDayOfMonth(time.Time(2019, 2, 7, 0, 0, 0, 0, time.UTC))  // Feb  7
=> time.Time(2019, 2, 28, 23, 59, 59, 0, time.UTC)           // Feb 28
*/
func LastDayOfMonth(t time.Time) time.Time {
	return FirstDayOfMonth(t).AddDate(0, 1, 0).Add(-time.Second)
}

/*
Takes a time.Time object and returns a time.Time object
which is the next day.

NextDay(time.Time(2019, 2, 7, 0, 0, 0, 0, time.UTC))   // Feb 7
=> time.Time(2019, 2, 8, 0, 0, 0, 0, time.UTC)         // Feb 8

NextDay(time.Time(2019, 2, 28, 0, 0, 0, 0, time.UTC))  // Feb 28
=> time.Time(2019, 3, 1, 0, 0, 0, 0, time.UTC)         // Mar  1
*/
func NextDay(t time.Time) time.Time {
	return t.AddDate(0, 0, 1)
}
