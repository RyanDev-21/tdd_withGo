package Geometry

import "testing"


func TestPerimeter(t *testing.T){
	rect :=Rectangle{10.0,10.0}
	got := rect.Perimeter()
	expected := 40.0

	if got!= expected{
		t.Errorf("got :%g want: %g",got, expected)
	}
}


func TestArea(t *testing.T){
		areaTests := []struct{
			name string
			shape Shape
			want float64
		}{
		{name:"Rectangle",shape:Rectangle{Width:12.0,Height:6.0},want:72.0},
		{name:"Circle",shape:Circle{Radius:10},want:314.1592653589793},
		{name:"Triangle",shape:Triangle{Base:12,Height:6},want:36.0},
		}
	
	for _,tt := range areaTests{
		t.Run(tt.name,func(t *testing.T) {

			got  := tt.shape.Area()
			if got !=tt.want{
				t.Errorf("got :%f want: %f",got, tt.want )
			}

	})		

	}



}
