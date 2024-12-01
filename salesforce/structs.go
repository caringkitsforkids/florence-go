/**
 *	Florence
 *	Copyright (C) 2024  hannjosh
 *
 *	This program is free software: you can redistribute it and/or modify
 *	it under the terms of the GNU General Public License as published by
 *	the Free Software Foundation, either version 3 of the License, or
 *	(at your option) any later version.
 *
 *	This program is distributed in the hope that it will be useful,
 *	but WITHOUT ANY WARRANTY; without even the implied warranty of
 *	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *	GNU General Public License for more details.
 *
 *	You should have received a copy of the GNU General Public License
 *	along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */
package salesforce

type Address struct {
	Street     string `json:"street"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postalCode"`
}

type Account struct {
	Name            string  `json:"Name"` // The name of the Account.
	ShippingAddress Address `json:"ShippingAddress"`
}

type Product2 struct {
	Name        string `json:"Name"`                  // The name of the Product.
	Description string `json:"Description,omitempty"` // Description of the Product.
	Family      string `json:"Family"`
}

type KitItem__c struct {
	Product  Product2 `json:"Product__r"`  // The Product in the kit.
	Quantity float32 `json:"Quantity__c"` // The Quantity (Number) of the Product in the kit. E.g. 2 x Long sleeve baby grow (0-3 months).
}
