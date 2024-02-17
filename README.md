# Introduction

As an example I tried implement advertising service of BV in this layout. Inside internal folder we had a three packages: **api**, **event_handler**, **lib**. </br>
- **api** - API service for CRUD operations of entities. Intended for dashboard users. (As a source used HTTP endpoints) </br>
- **event_handler** - Event handler service for managing advertising events coming from customers (like skip, watch, click events of ads). Also perform budget calculation of ads, according to received event. Intended for our customer services. (As a source used message queue) </br>
- **lib** - Library where now we store common command line arguments, which will be used in multiple services. </br>