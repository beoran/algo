#include <allegro5/allegro.h>

/* C Standard guarantees that all function pointers are equivalent. */
typedef void (*function_pointer)(void *);


int go_qsort_callback(void * p1, void * p2);

extern function_pointer go_upload_bitmap_cb;
extern function_pointer go_qsort_cb;

int take_callback(function_pointer cb, void * data);

typedef int go_take_callback_function(void * context);

int go_take_callback(void * context);

