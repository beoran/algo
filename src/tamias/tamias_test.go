package tamias

import "testing"


func assert(test * testing.T, cond bool, err string, args ... interface{})  {
  if cond {
    return
  }
  test.Errorf(err, args) 
}


func TestFloa(t * testing.T) {
  f1 := tamias.Float(10.0)
  f2 := tamias.Float(20.0)
  assert(t, f1 == 10.0, "Float should equal it's initial value 10.0", f1)
  assert(t, f1.Equals(10.0), "Float should equal it's initial value 10.0", f1)
  assert(t, f1.Min(f2) == f1, "Min must work properly.", f1)
  assert(t, f1.Max(f2) == f2, "Max must work properly.", f1)
  assert(t, f2.Min(f1) == f1, "Min must work properly in reverse.", f1)
  assert(t, f2.Max(f1) == f2, "Max must work properly in reverse.", f1)  
}

func TestBB(t * testing.T) {
  bb := tamias.BBMake(10.0, 20.0, 40.0, 80.0)  
  // assert(bb != nil , "Bounds Box must be constructable")
  assert(t, bb.L == 10.0, "Bounds Box L must be 10.0", bb.L)
  assert(t, bb.T == 20.0, "Bounds Box T must be 20.0", bb.T)
  assert(t, bb.R == 40.0, "Bounds Box R must be 40.0", bb.R)
  assert(t, bb.B == 80.0, "Bounds Box B must be 80.0", bb.B)
  b2 := bb.Grow(10.0)
  assert(t, b2.Contains(bb), "Contains and grow work correctly.", bb, b2)
}


func TestShape(t * testing.T) {
  body := tamias.BodyNew(10.0, 0.0)
  box  := tamias.BoxShapeNew(body, 20.0, 30.0)  
  box.CacheBB(body.Pos(), body.Rot())
  assert(t, box.GetBB() != nil, "Box must have a bounds box")
  if box.GetBB() != nil { 
    // the object should have been placed at (0,0), so half the BB
    // is positve and half negative
    // chipmunk, and hence Tamias too, use a normal Carthesian 
    // coordinate system where the zero is in the bottom left 
    assert(t, box.Shape.BB == box.GetBB(), "BB and GetBB() are the same")  
    assert(t, box.GetBB().L == -10.0, "Box must have BB.L -10.0", box.GetBB().L)
    assert(t, box.GetBB().T == 15.0, "Box must have BB.T -15.0", box.GetBB().T)
    assert(t, box.GetBB().R == 10.0, "Box must have BB.L 10.0", box.GetBB().R)
    assert(t, box.GetBB().B == -15.0, "Box must have BB.T -15.0", box.GetBB().B)
  }  
}

func TestVect(t * testing.T) {  
  v1 := tamias.VF(3.0, 4.0)
  v2 := tamias.V(1.0, 0.0)
  // tamias.V(3.0, 4.0)
  assert(t, v1.X == 3.0, "v1.X should be 3.0", v1.X)
  assert(t, v1.Y == 4.0, "v1.Y should be 4.0", v1.Y)
  assert(t, v2.X == 1.0, "v1.X should be 1.0", v2.X)
  assert(t, v2.Y == 0.0, "v1.Y should be 0.0", v2.Y)
  assert(t, v1.Length() == 5.0, "Vector length should be 5.")
  assert(t, v1.Equals(v1), "Vector should be equal to itself.")
  assert(t, v1.Add(v2).X == 4.0, "Vector Sum X should be 4.")
  assert(t, v1.Add(v2).Y == 4.0, "Vector Sum X should be 4.")
  vm :=	v1.Mult(tamias.Float(2.0))
  assert(t, vm.Y == 8.0, "Vector Mult Y should be 8.0.", vm.X, vm.Y)
  
  
}

func TestSpaceMap(t * testing.T) {  
  sm    := tamias.SpaceMapNew(10.0, 25)
  body  := tamias.BodyNew(10.0, 0.0)
  box   := tamias.BoxShapeNew(body, 20.0, 30.0)
  assert(sm != nil, "SpaceMap should be constructable")
  sm.Insert(box)
  bb    := box.GetBB().Grow(10.0)
  found := sm.FindBB(&bb)
  assert(t, found != nil, "SpaceMap should find back inserted items.", bb)
  if found != nil { 
    block := func(el interface {})(bool) {
      fmt.Println((el.(*tamias.PolyShape)))
      return el.(*tamias.PolyShape) == box
    } 
    res  := iterable.Find(found, block)
    assert(t, res != nil, "SpaceMap should find back the *right* inserted items.", found)
  } else {
    t.Errorf("SpaceMap could not find insertd item. (%s)\n", sm.String()) 
    // fmt.Printf()
  }  
   
  
}

