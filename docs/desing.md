# Design

## Tech stack
- be: go
    - db: postgres to cache data
    - go-etherium for intergation with blockchain
    - pino for signalling server WebRTC 
- fe: nextjs
    - WebRTC for video call
    - web3js
    - ethers.js
- smartcontract:
    - foundry
    - deploy testnet
- optional:
    - using SFU server for optimize stream ???

## Feature

- login with eth wallet
- create paid stream room
- user must pay fee to join room
- stream private with limited viewer
- donate for streamer
- notifications for donations

## API