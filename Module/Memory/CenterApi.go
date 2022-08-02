package center

func R_Memory_FromMap()map[int]string{
    return Memory_From_Map
}
func A_Memory_FromMap(name string){
    Memory_From_Map[len(Memory_From_Map)]=name
}
func S_Memory_FromMap(Set map[int]string){
    Memory_From_Map=Set
}