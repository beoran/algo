#include <allegro5/allegro.h>

/* C Standard guarantees that all function pointers are equivalent. */
typedef void (*function_pointer)(void);


bool go_upload_bitmap(void * bitmap, void *data);
extern function_pointer go_upload_bitmap_cb;

