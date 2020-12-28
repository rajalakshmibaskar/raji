package main
import "fmt"

func enqueue(queue[] int, element int) []int {
  queue = append(queue, element); 
  fmt.Println("Enqueued:", element);
  return queue
}

func dequeue(queue[] int) ([]int) {
  element := queue[0]; 
  fmt.Println("Dequeued:", element)
  return queue[1:]; 
}

func main() {
  var queue[] int; // Make a queue of ints.

  queue = enqueue(queue, 200);
  queue = enqueue(queue, 400);
  queue = enqueue(queue, 600);

  fmt.Println("Queue:", queue);

  queue = dequeue(queue);
  fmt.Println("Queue:", queue);

  queue = enqueue(queue, 800);
  fmt.Println("Queue:", queue);
}