#ifndef ALGO_CALLBACKS_H
#define ALGO_CALLBACKS_H


/* Callback for al_create_custom_bitmap */
ALLEGRO_BITMAP * go_create_custom_bitmap(int w, int h, void * context);

/* Callback for al_run_main_thread. */
int go_run_main_callback(int argc, char ** argv);

/* Callback for al_create_thread. */
void * go_create_thread_callback(ALLEGRO_THREAD * thread, void * arg);


#endif
