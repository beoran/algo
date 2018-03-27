#ifndef ALGO_CALLBACKS_H
#define ALGO_CALLBACKS_H


/* wrapper around al_create_custom_bitmap */
ALLEGRO_BITMAP * go_create_custom_bitmap(int w, int h, void * context);

int go_run_main_callback(int argc, char ** argv);


#endif
