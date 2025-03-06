import * as net from 'net'

// Start the server on port 64362
net
  .createServer(socket => {
    let buffer = Buffer.alloc(0)

    socket.on('data', data => {
      buffer = Buffer.concat([buffer, data])
      if (buffer.length < 2) return

      const size = parseInt(buffer.subarray(0, 2).toString())
      if (size <= 12 || isNaN(size)) {
        socket.end(`${13}ERR: incorrect size\n`)
        return
      }

      if (buffer.length < size + 2) return

      const request = buffer.subarray(2, size + 2)
      if (request.length <= 12) {
        socket.end(`${24}ERR: incorrect request length\n`)
        return
      }

      const cmd = request.subarray(4, 11).toString()
      if (cmd === 'RANDNUM') {
        const count = parseInt(request.subarray(12).toString())

        if (isNaN(count)) {
          socket.end(`${15}ERR: not a number\n`)
          return
        }

        if (count > 10) {
          socket.end(`${19}ERR: number too large\n`)
          return
        }

        const numbers = Array.from({length: count}, () =>
          Math.floor(Math.random() * 99)
        )
        const resp = numbers.join(',')
        const respLen =
          resp.length < 10 ? `0${resp.length}` : resp.length.toString()

        socket.end(`${respLen}${resp}\n`)
        return
      }

      socket.end(`${25}ERR: command not recognized\n`)
    })
  })
  .listen(64362, '127.0.0.1', () =>
    console.log('Server running on 127.0.0.1:64362')
  )