// configuration support
package al

/*
#include <stdlib.h>
#include <allegro5/allegro.h>
#include "helpers.h"
*/
import "C"
import "runtime"


type Config struct {
    handle * C.ALLEGRO_CONFIG
}

// Converts a config to it's underlying C pointer
func (self * Config) toC() *C.ALLEGRO_CONFIG {
    return (*C.ALLEGRO_CONFIG)(self.handle)
}

// Destroys the config.
func (self *Config) Destroy() {
    if self.handle != nil {
        C.al_destroy_config(self.toC())
    }
    self.handle = nil
}

// Wraps a C config into a go config
func wrapConfigRaw(data *C.ALLEGRO_CONFIG) *Config {
    if data == nil {
        return nil
    }
    return &Config{data}
}

// Sets up a finalizer for this Config that calls Destroy()
func (self *Config) SetDestroyFinalizer() *Config {
    if self != nil {
        runtime.SetFinalizer(self, func(me *Config) { me.Destroy() })
    }
    return self
}

// Wraps a C config into a go config and sets up a finalizer that calls Destroy()
func wrapConfig(data *C.ALLEGRO_CONFIG) *Config {
    self := wrapConfigRaw(data)
    return self.SetDestroyFinalizer()
}

type ConfigSection C.ALLEGRO_CONFIG_SECTION

func (cs * ConfigSection) toC() * C.ALLEGRO_CONFIG_SECTION {
    return (*C.ALLEGRO_CONFIG_SECTION)(cs)
}

func wrapConfigSectionRaw(ccs * C.ALLEGRO_CONFIG_SECTION) * ConfigSection {
    return (*ConfigSection)(ccs)
}

type ConfigEntry C.ALLEGRO_CONFIG_ENTRY

func (cs * ConfigEntry) toC() * C.ALLEGRO_CONFIG_ENTRY {
    return (*C.ALLEGRO_CONFIG_ENTRY)(cs)
}

func wrapConfigEntryRaw(ccs * C.ALLEGRO_CONFIG_ENTRY) * ConfigEntry {
    return (*ConfigEntry)(ccs)
}



func SystemConfig() * Config {
    return wrapConfig(C.al_get_system_config())
}


func CreateConfig() * Config {
    return wrapConfig(C.al_create_config())
}

func (cf * Config) AddSection(name string) {
    cname := cstr(name); defer cstrFree(cname)
    C.al_add_config_section(cf.toC(), cname) 
} 


func (cf * Config) SetValue(section, key, value string) {
    csection    := cstr(section);   defer cstrFree(csection)
    ckey        := cstr(key);       defer cstrFree(ckey)
    cvalue      := cstr(value);     defer cstrFree(cvalue)
    C.al_set_config_value(cf.toC(), csection, ckey, cvalue) 
} 

func (cf * Config) AddComment(section, key, comment string) {
    csection    := cstr(section);   defer cstrFree(csection)
    ckey        := cstr(key);       defer cstrFree(ckey)
    ccomment    := cstr(comment);   defer cstrFree(ccomment)
    C.al_set_config_value(cf.toC(), csection, ckey, ccomment) 
}

func (cf * Config) Value(section, key string) string {
    csection    := cstr(section);   defer cstrFree(csection)
    ckey        := cstr(key);       defer cstrFree(ckey)
    return C.GoString(C.al_get_config_value(cf.toC(), csection, ckey)) 
} 

func LoadConfig(filename string) * Config{
    cfilename := cstr(filename); defer cstrFree(cfilename)
    return wrapConfig(C.al_load_config_file(cfilename))
}

func LoadConfigFile(file * File) * Config{
    return wrapConfig(C.al_load_config_file_f(file.toC()))
}


func (cf * Config) Save(filename string) bool {
    cfilename := cstr(filename); defer cstrFree(cfilename)
    return bool(C.al_save_config_file(cfilename, cf.toC()))
}

func (cf * Config) SaveFile(file * File) bool {
    return bool(C.al_save_config_file_f(file.toC(), cf.toC()))
}

func (cf * Config) MergeInto(add * Config) {
    C.al_merge_config_into(cf.toC(), add.toC())
}

func (cf * Config) Merge(cf2 * Config) * Config {
    return wrapConfig(C.al_merge_config(cf.toC(), cf2.toC()))
}


func (cf * Config) RemoveSection(name string) bool {
    cname := cstr(name); defer cstrFree(cname)
    return bool(C.al_remove_config_section(cf.toC(), cname)) 
} 


func (cf * Config) RemoveKey(section, key string) bool {
    csection    := cstr(section);   defer cstrFree(csection)
    ckey        := cstr(key);       defer cstrFree(ckey)
    return bool(C.al_remove_config_key(cf.toC(), csection, ckey))
} 

func (cf * Config) FirstSection() (name string, section * ConfigSection) {
    section   = &ConfigSection{}
    csection := section.toC()
    name = C.GoString(C.al_get_first_config_section(cf.toC(), &csection))
    return name, section
}

func (section * ConfigSection) Next() (name string, ok bool) {
    csection := section.toC()
    cname := C.al_get_next_config_section(&csection)
    if cname == nil {
        name    = ""
        ok      = false
    } else {
        name    = C.GoString(cname)
        ok      = true
    }
    return name, ok
}

func (cf * Config) FirstEntry(section string) (name string, entry * ConfigEntry) {
    csection := cstr(section);   defer cstrFree(csection)
    entry   = &ConfigEntry{}
    centry := entry.toC()
    name = C.GoString(C.al_get_first_config_entry(cf.toC(), csection, &centry))
    return name, entry
}

func (entry * ConfigEntry) Next() (name string, ok bool) {
    centry := entry.toC()
    cname  := C.al_get_next_config_entry(&centry)
    if cname == nil {
        name    = ""
        ok      = false
    } else {
        name    = C.GoString(cname)
        ok      = true
    }
    return name, ok
}


