package routes

import (
	"TakeHomeApi/pkg/schemas"
	"TakeHomeApi/pkg/stores"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Reset(store *stores.AccountStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		store.Reset()
		c.String(http.StatusOK, "OK")
	}
}

func Balance(account *stores.AccountStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		accountID := c.Query("account_id")

		balance,exists := account.GetBalance(accountID)

		if !exists {
			c.String(http.StatusNotFound, "0")
			return
		}
		c.String(http.StatusOK, strconv.Itoa(balance))
	}
}


func Event(account *stores.AccountStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		
		var event schemas.RequestEvent
		
		if err := c.ShouldBindJSON(&event); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		
		if event.Type == "deposit" {
			if event.Destination == "" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Destination required for deposit"})
				return
			}

			account.AddBalance(event.Destination, event.Amount)
			balance,_ := account.GetBalance(event.Destination)

			response := schemas.EventResponse{
				Destination: &schemas.Account{ID: event.Destination, Balance: balance},
			}
			c.JSON(http.StatusCreated, response)

		} else if event.Type == "withdraw" {
			if event.Origin == "" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Origin required for withdraw"})
				return
			}
			
			balance, exists := account.GetBalance(event.Origin)

			if exists{
				
				if balance >= event.Amount {

					account.SubtractBalance(event.Origin, event.Amount)

					updatedBalance,_ := account.GetBalance(event.Origin)
					
					response := schemas.EventResponse{
						Origin: &schemas.Account{ID: event.Origin, Balance: updatedBalance},
					}

					c.JSON(http.StatusCreated, response)
				} else {
					c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient funds"})
				}

			} else {
				c.String(http.StatusNotFound, "0")
			}

		} else if event.Type == "transfer" {
			if event.Origin == "" || event.Destination == "" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Origin and destination required for transfer"})
				return
			}
			
			originBalance, originExists := account.GetBalance(event.Origin)

			// The the destination is not being checked in the transfer
			// it will be created anyway
			if originExists {
				if originBalance >= event.Amount {
					
					account.SubtractBalance(event.Origin, event.Amount)
					account.AddBalance(event.Destination, event.Amount)
					updatedOriginBalance,_ := account.GetBalance(event.Origin)
					destinationBalance,_ := account.GetBalance(event.Destination)

					response := schemas.EventResponse{
						Origin:      &schemas.Account{ID: event.Origin, Balance: updatedOriginBalance},
						Destination: &schemas.Account{ID: event.Destination, Balance: destinationBalance},
					}
					c.JSON(http.StatusCreated, response)
				} else {
					c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient funds"})
				}
			} else {
				c.String(http.StatusNotFound, "0")
			}

		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event type"})
		}

	}
}
