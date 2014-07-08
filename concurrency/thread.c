#include <stdio.h>
#include <unistd.h>
#include <pthread.h>

void* task(void* argument) {
  // Do a bunch of stuff.
  return (NULL);
}

int main() { 
  pthread_t thread;
  int result;

  result = pthread_create(&thread, NULL, task, NULL); // HL

  if (result != 0) {
	 fprintf(stderr, "error creating thread (errno = %d)\n",result);
	 exit(0);
  }

  pthread_join(thread, NULL);  // HL
}
