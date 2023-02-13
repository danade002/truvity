#Here is the architecture of the system represented in UML (Unified Modeling Language):


@startuml
title System Architecture

package "SPA Frontend" {
  [SPA Frontend]
}

package "API Backend" {
  [API Backend]
}

database "Postgres Cluster" {
  [Postgres Cluster]
}

cloud "S3 Bucket" {
  [S3 Bucket]
}

cloud "External Data Provider" {
  [External Data Provider]
}

database "SQL Script for Initializing the Database" {
  [SQL Script for Initializing the Database]
}

package "Binary for Creating Fixtures" {
  [Binary for Creating Fixtures]
}

[SPA Frontend] --> [API Backend]
[API Backend] --> [Postgres Cluster]
[API Backend] --> [S3 Bucket]
[API Backend] --> [External Data Provider]
[SQL Script for Initializing the Database] --> [Postgres Cluster]
[Binary for Creating Fixtures] --> [Postgres Cluster]

@enduml
