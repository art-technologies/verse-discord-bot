# Verse Discord Bot

- [Introduction](#introduction)
- [Add it to your Discord server](#add-it-to-your-discord-server)
- [Configuration](#configuration)
  - [Filters](#filters)
    - [Artists filter](#artists-filter)
    - [Collaborators filter](#collaborators-filter)
    - [Collcetions filter](#collcetions-filter)
    - [Events filter](#events-filter)
  - [Configuration examples](#configuration-examples)
  - [Example 1](#example-1)
  - [Example 2](#example-2)
  - [Example 3](#example-3)

<small><i><a href='http://ecotrust-canada.github.io/markdown-toc/'>Table of contents generated with markdown-toc</a></i></small>

## Introduction

Verse Discord bot allows to stream real time marketplace events to your Discord server. It can be configured to send events for sales, listings and offers including both Verse and OpenSea.

Note: OpenSea events will work only for projects that exist on Verse.

## Add it to your Discord server

In order to start using Verse discord bot you should

1. [Invite Verse Discord Bot](https://discord.com/oauth2/authorize?client_id=1180153976331706408&permissions=2048&scope=bot) to your server and create a channel for the events.
2. Identify what events you want to show in your Discord server and create [configuration](#Configuration).
3. Make Pull Request to [guilds.json](./cmd/server/guilds.json) with your configuration requirements based. If you are can't make Pull Request please reach out to us at support@verse.works.
4. Once we approve Pull Request your Discord channel will start receiving events.

## Configuration

Bot configuration describes what sale/offer/lising events to send and where to send them. The most general bot configuration is the following and consists of `channel_id` and optional filters for `collections`, `artists` and `collaborators`.

```json
{
  "channel_id": "123",
  "filters": {
    "collections": ["33-million-by-anna-lucia"],
    "collaborators": ["tender", "verse-solos"],
    "artists": ["anna-lucia"],
    "events": [
      "PM_SALE",
      "SM_LISTED",
      "SM_SALE",
      "SM_OFFER",
      "SM_GLOBAL_OFFER",
      "OS_SALE",
      "OS_OFFER",
      "OS_LISTED"
    ]
  }
}
```

You get your Discord channel ID by right clicking on the channel and "Copy Channel ID".

### Filters

If you don't want to receive all the events you can specify optional filters to receive only the events you want.

#### Artists filter

Allows you to receive events only for one or more specific artists.

#### Collaborators filter

Allows you to reeive events only for one or more specific collaborators/galleries.

#### Collcetions filter

Allows you to receive events only for one or more specific collections.

#### Events filter

Allows to specify what kind of sale, offer or listing events should be sent.

| Event             | Description                             |
| ----------------- | --------------------------------------- |
| `PM_SALE`         | Verse primary market sale events        |
| `SM_LISTED`       | Verse secondary market listing events   |
| `SM_SALE`         | Verse secondary market sale events      |
| `SM_OFFER`        | Verse offer events                      |
| `SM_GLOBAL_OFFER` | Verse secondary market offer events     |
| `OS_SALE`         | OpenSea secondary market sale events    |
| `OS_OFFER `       | OpenSea offer events                    |
| `OS_LISTED`       | OpenSea secondary market listing events |

### Configuration examples

Examples for commonly used filters

### Example 1

All events (sale, listing, offer) for [Tender](https://verse.works/tender) gallery.

```json
{
  "channel_id": "123",
  "filters": {
    "collaborators": ["tender"]
  }
}
```

### Example 2

All events (sale, listing, offer) for artist [qubibi](https://verse.works/qubibi).

```json
{
  "channel_id": "123",
  "filters": {
    "artists": ["qubibi"]
  }
}
```

### Example 3

Sale events for [Quasi Dragon Studies](https://verse.works/series/quasi-dragon-studies-by-harvey-rayner) collection.

```json
{
  "channel_id": "123",
  "filters": {
    "collections": ["quasi-dragon-studies-by-harvey-rayner"],
    "events": ["PM_SALE", "SM_SALE", "OS_SALE"]
  }
}
```
