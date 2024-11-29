const add = (a: number, b: number) => a + b

const mult = (a: number, b: number) => a * b

for (let i = 0; i < 10; i++) {
  process.stdout.write(Buffer.from(`from ts script: ${i}\n`))
}