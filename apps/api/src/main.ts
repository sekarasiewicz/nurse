import * as express from 'express'
import { Message } from '@nurse/api-interfaces'
import { Request, Response } from "express"
import { createConnection } from "typeorm"
import { User } from './entity'


createConnection().then(connection => {
  const userRepository = connection.getRepository(User)

  // create and setup express app
  const app = express()
  app.use(express.json())

  // register routes

  app.get("/users", async function(req: Request, res: Response) {
    const users = await userRepository.find()
    res.json(users)
  })

  app.get("/users/:id", async function(req: Request, res: Response) {
    const results = await userRepository.findOne(req.params.id)
    return res.send(results)
  })

  app.post("/users", async function(req: Request, res: Response) {
    const user = userRepository.create(req.body)
    const results = await userRepository.save(user)
    return res.send(results)
  })

  app.put("/users/:id", async function(req: Request, res: Response) {
    const user = await userRepository.findOne(req.params.id)
    userRepository.merge(user, req.body)
    const results = await userRepository.save(user)
    return res.send(results)
  })

  app.delete("/users/:id", async function(req: Request, res: Response) {
    const results = await userRepository.delete(req.params.id)
    return res.send(results)
  })

  const port = process.env.port || 3333
  const server = app.listen(port, () => {
    console.log('Listening at http://localhost:' + port + '/api')
  })
  server.on('error', console.error)
})

// const app = express()
// app.use(express.json())
//
// const greeting: Message = { message: 'Welcome to api!' }
//
// app.get('/api', (req, res) => {
//   res.send(greeting)
// })
//
// // register routes
//
// app.get("/users", (req: Request, res: Response) => {
//   // here we will have logic to return all users
// })
//
// app.get("/users/:id", (req: Request, res: Response) => {
//   // here we will have logic to return user by id
// })
//
// app.post("/users", (req: Request, res: Response) => {
//   // here we will have logic to save a user
// })
//
// app.put("/users/:id", (req: Request, res: Response) => {
//   // here we will have logic to update a user by a given user id
// })
//
// app.delete("/users/:id", (req: Request, res: Response) => {
//   // here we will have logic to delete a user by a given user id
// })
//
// const port = process.env.port || 3333
// const server = app.listen(port, () => {
//   console.log('Listening at http://localhost:' + port + '/api')
// })
// server.on('error', console.error)
