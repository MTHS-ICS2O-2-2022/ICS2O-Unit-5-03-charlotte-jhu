// Copyright (c) 2020 Mr. Coxall All rights reserved
//
// Created by: Charlotte Jhu
// Created on: May 2023
// This file contains the JS functions for index.html

'use strict'

function myButtonClicked() {
// This function tells the user what rated movies they can watch
// input
  const age = parseInt(document.getElementById("userAge").value)

// process
  if (age >= 17) {
    // output
    document.getElementById("answer").innerHTML = "You can watch rated R movies alone!"
  } else if (age >= 13) {
    // output
    document.getElementById("answer").innerHTML = "You can watch rated PG-13 movies alone!"
  } else if (age >= 5) {
    // output
    document.getElementById("answer").innerHTML = "You can watch rated G or PG movies alone!"
  } else {
    // output
    document.getElementById("answer").innerHTML = "You are too young to watch any movies alone!"
  }
}
