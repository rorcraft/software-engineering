// You are given a list of items / combo_items with their price list.
// And you are given list of items to buy.
// Now you are asked to find which combination to buy so that it costs you minimum.
// It doesnt matter if you are getting some extra items if it costs less.
//
//
// Sr.No Price | Items/Combo_Items
// 1.	5	|	Burger
// 2.	4	|	French_Frice
// 3.	8	|	Coldrink
// 4.	12	|	Burger, French_Frice, Coldrink
// 5.	14 | Burger, Coldrink
//
//
// Input Items to Buy:
// Coldrink
//
// Output(Sr.No)
// 3
//
// Input Items to Buy:
// Burger Coldrink
//
// Output(Sr.No)
// 4
//
// Input Items to Buy:
// Burger French_Frice
//
// Output(Sr.No)
// 1,2
var _ = require('lodash')
var combo = [
  [1, 5, ['Burger']]
, [2, 4, ['French_Fries', 'Ice_Cream']]
, [3, 1, ['Colddrink', 'Ice_Cream']]
, [4, 16, ['Burger', 'French_Fries', 'Colddrink']]
, [5, 14, ['Burger', 'Colddrink']]
]

function indexItem(combo) {
  var items = {}
  combo.forEach(function(lineItem) {
    // console.log(lineItem)
    lineItem[2].forEach(function (item) {
      items[item] = items[item] || []
      items[item].push(lineItem.slice(0,2))
    })
  })
  return items
}

var nameToPrice = indexItem(combo)

function buy(items) {
  var prices = []
    , priceCombos = []
    , cheapestKeys = []
  items.forEach(function (item) {
    prices.push(nameToPrice[item])
  })
  // console.log(prices)
  priceCombos = pricesToCombos(prices, priceCombos)
  // console.log(priceCombos)
  cheapest = findCheapestCombo(priceCombos)
  cheapest.forEach(function (keyPrice) {
    cheapestKeys.push(keyPrice[0])
  })
  console.log(cheapestKeys)
  return cheapestKeys
}

function pricesToCombos(prices, priceCombos) {
  if (prices.length == 0) return priceCombos
  var currentPrices = prices.shift()
    , newPriceCombos = []
  if (priceCombos.length == 0) {
    currentPrices.forEach(function (keyPrice) {
      newPriceCombos.push([keyPrice])
    })
  } else {
    priceCombos.forEach(function (priceCombo) {
      currentPrices.forEach(function (keyPrice) {
        newPriceCombos.push(priceCombo.concat([keyPrice]))
      })
    })
  }
  return pricesToCombos(prices, newPriceCombos)
}

function findCheapestCombo(priceCombos) {
  var cheapestComboTotal = Number.MAX_VALUE,
      cheapestCombo = []
  priceCombos.forEach(function (combo) {
    combo = _.uniq(combo, function (keyPrice) { return keyPrice[0] })
    // console.log(combo)
    var total = 0
    _.each(combo, function (keyPrice) {
      total += keyPrice[1]
    })
    // console.log(total)
    if (total < cheapestComboTotal) {
      cheapestComboTotal = total
      cheapestCombo = combo
    }
  })
  return cheapestCombo
}

buy(['Burger', 'Colddrink', 'French_Fries'])
// => [ 1, 3, 2 ]
buy(['Burger'])
// => [1]
buy(['Colddrink'])
// => [3]
buy(['French_Fries', 'Ice_Cream'])
// => [2]
