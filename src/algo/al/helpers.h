#include <allegro5/allegro.h>

bool go_upload_bitmap(void * bitmap, void *data);
typedef bool go_upload_bitmap_function(void  * bitmap, void *data);
extern go_upload_bitmap_function * go_upload_bitmap_cb;

