type MyHashSet struct {
    Data []bool
}

func Constructor() MyHashSet {
  return MyHashSet{Data : make([]bool, 1000001)}
}

func (this *MyHashSet) Add(key int) {
    this.Data[key] = true
}

func (this *MyHashSet) Remove(key int) {
   this.Data[key] = false
}

func (this *MyHashSet) Contains(key int) bool {
    return this.Data[key]    
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
 