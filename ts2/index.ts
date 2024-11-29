const add = (a: number, b: number) => a + b

const mult = (a: number, b: number) => a * b

process.stdin.on('data', data => 
  data.toString()
    .split("\n")
    .filter(s => s.length)
    .forEach(s => console.log(`from ts2 script: ${s}`)))