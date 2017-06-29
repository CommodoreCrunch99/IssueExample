package main

import (
	"azul3d.org/engine/gfx"
	"azul3d.org/engine/gfx/gfxutil"
	"gridworldapp/tools"
	"path"
	"log"
	"io"
)

func CreateWallObject() *gfx.Object{
	squareMesh := gfx.NewMesh()
	sideLength := float32(1.0)
	squareMesh.Vertices = []gfx.Vec3{
		//Top
		{sideLength/2.0,0,sideLength/2.0},
		{-sideLength/2.0,0,sideLength/2.0},
		{sideLength/2.0,0,-sideLength/2.0},
		//Bottom
		{sideLength/2.0,0,-sideLength/2.0},
		{-sideLength/2.0,0,sideLength/2.0},
		{-sideLength/2.0,0,-sideLength/2.0},
	}

	squareMesh.Colors = []gfx.Color{
		//Top
		{0,0,0,1},
		{0,0,0,1},
		{0,0,0,1},
		//Bottom
		{0,0,0,1},
		{0,0,0,1},
		{0,0,0,1},
	}

	squareObj := gfx.NewObject()
	shader, err := gfxutil.OpenShader(tools.Path(path.Join("shaders","basic"))) //TODO
	if err != nil {
		log.Print(err)
	}
	squareObj.Shader = shader
	squareObj.OcclusionTest = true
	squareObj.State = gfx.NewState()
	squareObj.AlphaMode = gfx.NoAlpha
	squareObj.State.FaceCulling = gfx.BackFaceCulling
	squareObj.Meshes = []*gfx.Mesh{squareMesh}

	return squareObj
}

type qColors []*gfx.Color //assume 0 is top, 1 bottom, 2 left, 3 right

type qValSquare *gfx.Object

func CreateQValSquare() *qValSquare{
	squareMesh := gfx.NewMesh()
	sideLength := float32(1.0)
	squareMesh.Vertices = []gfx.Vec3{
		//Top Triangle
		{sideLength/2.0,0,sideLength/2.0}, //top right
		{-sideLength/2.0,0,sideLength/2.0},//top Left
		{0,0,0},//center
		//Bottom Triangle
		{sideLength/2.0,0,-sideLength/2.0}, //Bottom right
		{-sideLength/2.0,0,-sideLength/2.0},//Bottom Left
		{0,0,0},//center
		//Left Triangle
		{-sideLength/2.0,0,sideLength/2.0}, //top Left
		{-sideLength/2.0,0,-sideLength/2.0},//Bottom Left
		{0,0,0},//center
		//Right Triangle
		{sideLength/2.0,0,sideLength/2.0}, //top Left
		{sideLength/2.0,0,-sideLength/2.0},//Bottom Left
		{0,0,0},//center
	}

	squareMesh.Colors = []gfx.Color{
		//Top
		{0,0,.5,1},
		{0,0,.5,1},
		{0,0,.5,1},
		//Bottom
		{0,0,.5,1},
		{0,0,.5,1},
		{0,0,.5,1},
		//Left
		{0,0,.5,1},
		{0,0,.5,1},
		{0,0,.5,1},
		//Right
		{0,0,.5,1},
		{0,0,.5,1},
		{0,0,.5,1},
	}

	squareObj := gfx.NewObject()

	shader, err := gfxutil.OpenShader(tools.Path(path.Join("shaders","basic"))) //TODO
	if err != nil {
		log.Print(err)
	}
	squareObj.Shader = shader
	squareObj.OcclusionTest = true
	squareObj.State = gfx.NewState()
	squareObj.State.FaceCulling = gfx.BackFaceCulling
	squareObj.Meshes = []*gfx.Mesh{squareMesh}

	return &qValSquare(*squareObj)
}

func changeQColor(qv qValSquare, c qColors) {
	avColor := averageColor(c...)
	for key, value := range qv.Meshes[0].Colors {
		if (key+1)%3 != 0{
			value.R = c[key/3].R
			value.G = c[key/3].G
			value.B = c[key/3].B
			value.A = c[key/3].A
		} else {
			value.R = avColor.R
			value.G = avColor.G
			value.B = avColor.B
			value.A = avColor.A
		}
	}
}

func averageColor(colors ...*gfx.Color) *gfx.Color{
	avg := gfx.Color{}
	for key, value := range colors{
		if key == 0{
			avg.R = value.R
			avg.G = value.G
			avg.B = value.B
			avg.A = value.A
		}else{
			avg.R += value.R
			avg.R /= 2.0
			avg.G += value.G
			avg.G /= 2.0
			avg.B += value.B
			avg.B /= 2.0
			avg.A += value.A
			avg.A /= 2.0
		}
	}
	return &avg
}





