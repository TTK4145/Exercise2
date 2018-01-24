#include <pthread.h>
#include <stdio.h>

int i = 0;
pthread_mutex_t mutex = PTHREAD_MUTEX_INITIALIZER;
//pthread_mutex_init(&mutex, NULL);


// Note the return type: void*
void* incrementingThreadFunction(){
	int j;
    for (j = 0; j < 1000000; j++) {
	// TODO: sync access to i
	pthread_mutex_lock(&mutex);
	i++;
	pthread_mutex_unlock(&mutex);
    }
    return NULL;
}

void* decrementingThreadFunction(){
	int j;
    for (j = 0; j < 1000001; j++) {
	// TODO: sync access to i
	pthread_mutex_lock(&mutex);
	i--;
	pthread_mutex_unlock(&mutex);
    }
    return NULL;
}


int main(){
    pthread_t incrementingThread, decrementingThread;
    
    pthread_create(&incrementingThread, NULL, incrementingThreadFunction, NULL);
    pthread_create(&decrementingThread, NULL, decrementingThreadFunction, NULL);
    
    pthread_join(incrementingThread, NULL);
    pthread_join(decrementingThread, NULL);
    pthread_mutex_destroy(&mutex);
    printf("The magic number is: %d\n", i);
    return 0;
}
