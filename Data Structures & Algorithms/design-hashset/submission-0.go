type MyHashSet struct {
    Data []int
}

func Constructor() MyHashSet {

  return MyHashSet{Data : []int{}}
    
}

func (this *MyHashSet) Add(key int) {
    if !this.Contains(key){
        this.Data = append(this.Data, key)
    }
}

func (this *MyHashSet) Remove(key int) {
    for i, k := range this.Data {
        if k == key {
            this.Data = append(this.Data[:i], this.Data[i+1:]...)
        }
    }
    
}

func (this *MyHashSet) Contains(key int) bool {

    for _, k := range this.Data {
        if k == key {
            return true
        }
    }

    return false
    
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
 