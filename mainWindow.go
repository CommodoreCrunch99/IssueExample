package main

import (
	"azul3d.org/engine/gfx/window"
	"log"
	"time"
	"math/rand"
	"os"
	"gridworldapp/tools"
	"azul3d.org/engine/gfx"
	"azul3d.org/engine/gfx/camera"
	"azul3d.org/engine/gfx/gfxutil"
	azulmath "azul3d.org/engine/lmath"
	"path"
	"fmt"
)

var shaderDir string

func createWalls(num int, shader *gfx.Shader , color gfx.Color) (output []*gfx.Object){
	output = make([]*gfx.Object,0)

	for len(output) < num{
		triMesh:=gfx.NewMesh()
		triMesh.Vertices = []gfx.Vec3{//square mesh
			{1,0,1},
			{-1,0,1},
			{1,0,-1},

			{1,0,-1},
			{-1,0,1},
			{-1,0,-1},
		}
		triMesh.Colors = []gfx.Color{
			color,
			color,
			color,

			color,
			color,
			color,
		}

		triObj := gfx.NewObject()
		triObj.Shader = shader
		triObj.OcclusionTest = true
		triObj.State = gfx.NewState()
		triObj.State.FaceCulling = gfx.NoFaceCulling
		triObj.Meshes = []*gfx.Mesh{triMesh}

		//append object to slice
		output = append(output, triObj)
	}
	return
}

func createGoal(shader *gfx.Shader, color gfx.Color) (output *gfx.Object){
	triMesh:=gfx.NewMesh()
	triMesh.Vertices = []gfx.Vec3{//square mesh
		{1,0,1},
		{-1,0,1},
		{1,0,-1},

		{1,0,-1},
		{-1,0,1},
		{-1,0,-1},
	}
	triMesh.Colors = []gfx.Color{
		color,
		color,
		color,

		color,
		color,
		color,
	}

	output = gfx.NewObject()
	output.Shader = shader
	output.OcclusionTest = true
	output.State = gfx.NewState()
	output.State.FaceCulling = gfx.NoFaceCulling
	output.Meshes = []*gfx.Mesh{triMesh}

	return
}

func createAgent(shader *gfx.Shader, color gfx.Color) (output *gfx.Object){
	triMesh:=gfx.NewMesh()
	triMesh.Vertices = []gfx.Vec3{//square mesh
		{1,0,1},
		{-1,0,1},
		{1,0,-1},

		{1,0,-1},
		{-1,0,1},
		{-1,0,-1},
	}
	triMesh.Colors = []gfx.Color{
		color,
		color,
		color,

		color,
		color,
		color,
	}

	output = gfx.NewObject()
	output.Shader = shader
	output.OcclusionTest = true
	output.State = gfx.NewState()
	output.State.FaceCulling = gfx.NoFaceCulling
	output.Meshes = []*gfx.Mesh{triMesh}

	return
}

func gfxLoop(w window.Window, d gfx.Device,){

	//create orthographic camera
	mainCamera := camera.NewOrtho(d.Bounds())
	mainCamera.SetPos(azulmath.Vec3{0,-2,0})

	//read shader from file
	shader, err := gfxutil.OpenShader(path.Join(shaderDir, "basic"))
	if err != nil {
		fmt.Println(err)
		time.Sleep(time.Duration(10) * time.Second)
		log.Fatal(err)
	}

	//create object List to hold all objects to be rendered
	renderList := make([]*gfx.Object,0)

	//create walls
	wallPool := createWalls(10,shader,gfx.Color{1,1,1,1})
	//attach walls
	renderList = append(renderList,wallPool...)

	//Create Goal
	goal := createGoal(shader, gfx.Color{1,.75,0,1})
	//attach goal
	renderList = append(renderList,goal)

	//Create Agent
	agent := createAgent(shader, gfx.Color{.25,0,1,1})
	//attach agent
	renderList = append(renderList,agent)

	//render loop
	for{
		//clear screen, worry about optimizing later
		d.Clear(d.Bounds(), gfx.Color{1, 1, 1, 1})//clear background to white
		d.ClearDepth(d.Bounds(),1)

		//redraw all objects
		for _,object := range renderList{
			d.Draw(d.Bounds(),object, mainCamera)
		}
		//final render
		d.Render()
	}
}

func main() {
	//goPool := sync.WaitGroup{}
	//initialization & setup
	rand.Seed(time.Now().UnixNano())

	logFile, err := os.Create((time.Now().Format("2_JAN_1504")+".log"))
	if err!=nil{
		log.Fatalln(err)//prevent program from being unable to log errors
	}
	log.SetOutput(logFile)//set log output to file

	shaderDir = tools.Path("shaders")

	//go startAI()// start ai thread


	// Create a window and run our graphics loop.
	window.Run(gfxLoop, nil)
}