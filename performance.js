function intToAlphabet(num = 0) {
  if (num === 0) {
    return 'a'
  }
  const ALPHA = 'abcdefghijklmnopqrstuvwxyz'
  let res = ''
  while (num) {
    res = ALPHA.charAt(num % ALPHA.length) + res
    num = Math.floor(num / ALPHA.length)
  }
  return res
}

function mapOperation(count) {
  if(!count) count = 3000000
  const dict = {}
  for (let i = 0; i < count; i++) {
    dict[intToAlphabet(i)] = i
  }
  let sum = 0
  for (let key in dict) {
    sum += dict[key]
  }
  return sum
}

function test(operation, count, args) {
  const t0 = new Date().getTime()
  for (let i = 0; i < count; i++) {
    operation(args);
    console.log(`[${operation.name}] Round ${i} done.`)
  }
  console.log(`[${operation.name}] Time used: ${(new Date().getTime() - t0) / 1000}s`)
}

test(mapOperation, 50)
