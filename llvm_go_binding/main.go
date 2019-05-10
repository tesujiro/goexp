package main

import (
	"fmt"

	"llvm.org/llvm/bindings/go/llvm"
)

func main() {
	// Initialize the MCJIT
	llvm.LinkInMCJIT()
	llvm.InitializeNativeTarget()
	llvm.InitializeNativeAsmPrinter()
	llvm.InitializeAllAsmParsers()

	// setup our builder and module
	builder := llvm.NewBuilder()
	mod := llvm.NewModule("my_module")

	// Create main function
	main := llvm.FunctionType(llvm.Int32Type(), []llvm.Type{}, false)
	llvm.AddFunction(mod, "main", main)
	block := llvm.AddBasicBlock(mod.NamedFunction("main"), "entry")
	builder.SetInsertPoint(block, block.FirstInstruction())

	// Declare the puts
	i8p := llvm.PointerType(llvm.Int8Type(), 0)
	ftPuts := llvm.FunctionType(llvm.Int32Type(), []llvm.Type{i8p}, false)
	puts := llvm.AddFunction(mod, "puts", ftPuts)

	// Declare the memcpy
	ftMemcpy := llvm.FunctionType(llvm.VoidType(), []llvm.Type{
		llvm.PointerType(llvm.IntType(8), 0),
		llvm.PointerType(llvm.IntType(8), 0),
		llvm.IntType(32),
		//llvm.IntType(32),
		llvm.IntType(1),
	}, false)
	memcpy := llvm.AddFunction(mod, "llvm.memcpy.p0i8.p0i8.i32", ftMemcpy)

	// Allocate Array
	ar := builder.CreateArrayAlloca(llvm.Int8Type(), llvm.ConstInt(llvm.Int32Type(), 15, false), "ar")

	// Get Pointer *i8
	sp := builder.CreateGEP(ar, []llvm.Value{llvm.ConstInt(llvm.Int64Type(), 0, false)}, "sp")
	// sp := builder.CreateBitCast(ar, llvm.PointerType(llvm.Int8Type(), 0), "sp")
	s := builder.CreateGlobalStringPtr("hogehogeaiueo", "s")

	// Call the memcpy
	builder.CreateCall(memcpy, []llvm.Value{
		sp, s, llvm.ConstInt(llvm.Int32Type(), 15, false),
		//llvm.ConstInt(llvm.IntType(32), 1, false),
		llvm.ConstInt(llvm.IntType(1), 1, false),
	}, "")

	// Call the puts
	builder.CreateCall(puts, []llvm.Value{sp}, "")

	// Return Code
	builder.CreateRet(llvm.ConstInt(llvm.IntType(32), 0, false))

	// Verify Module
	if err := llvm.VerifyModule(mod, llvm.ReturnStatusAction); err != nil {
		fmt.Println(err.Error())
		return
	}

	// Set MCJITOption
	opt := llvm.NewMCJITCompilerOptions()
	opt.SetMCJITEnableFastISel(true)
	opt.SetMCJITOptimizationLevel(2)
	opt.SetMCJITNoFramePointerElim(true)
	opt.SetMCJITCodeModel(llvm.CodeModelJITDefault)

	// Make MCJIT
	engine, err := llvm.NewMCJITCompiler(mod, opt)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer engine.Dispose()

	// Optimazetion
	pass := llvm.NewPassManager()
	defer pass.Dispose()

	/*
		pass.Add(engine.TargetData())
		pass.AddConstantPropagationPass()
		pass.AddInstructionCombiningPass()
		pass.AddPromoteMemoryToRegisterPass()
		pass.AddGVNPass()
		pass.AddCFGSimplificationPass()
	*/
	pass.Run(mod)

	// Output llvm IR
	mod.Dump()

	// Run the function
	funcResult := engine.RunFunction(mod.NamedFunction("main"), []llvm.GenericValue{})
	fmt.Printf("%d\n", funcResult.Int(false))
}
