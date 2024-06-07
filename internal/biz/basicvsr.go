package biz

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sbabiv/xml2map"
	"os/exec"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
)

// BasicVSR is a BasicVSR model.
type BasicVSR struct {
	ProductName   string `json:"product_name"`
	DriverVersion string `json:"driver_version"`
	CudaVersion   string `json:"cuda_version"`
	FanSpeed      string `json:"fan_speed"`

	FbMemoryUsage FbMemoryUsage `json:"fb_memory_usage"`
	Temperature   Temperature   `json:"temperature"`
	PowerReadings PowerReadings `json:"power_readings"`
	Processes     Processes     `json:"processes"`
}

type FbMemoryUsage struct {
	Total    string `json:"total"`
	Reserved string `json:"reserved"`
	Used     string `json:"used"`
	Free     string `json:"free"`
}

type Temperature struct {
	GpuTemp string `json:"gpu_temp"`
}

type PowerReadings struct {
	PowerDraw string `json:"power_draw"`
}

type Processes struct {
	ProcessInfo []ProcessInfo `json:"process_info"`
}

type ProcessInfo struct {
	Pid         string `json:"pid"`
	ProcessName string `json:"process_name"`
	UsedMemory  string `json:"used_memory"`
}

type GPURequest struct {
	Name string `json:"name"`
}

type ExecResult struct {
	Name string `json:"name"`
}

// BasicVSRRepo is a BasicVSRRepo repo.
type BasicVSRRepo interface {
	Save(context.Context, *BasicVSR) (*BasicVSR, error)
}

// BasicVSRUsecase is a BasicVSR usecase.
type BasicVSRUsecase struct {
	repo BasicVSRRepo
	log  *log.Helper
}

// NewBasicVSRUsecase new a BasicVSR usecase.
func NewBasicVSRUsecase(repo BasicVSRRepo, logger log.Logger) *BasicVSRUsecase {
	return &BasicVSRUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateBasicVSR creates a BasicVSR, and returns the new BasicVSR.
func (uc *BasicVSRUsecase) GetStatus(ctx context.Context, g *BasicVSR) (*BasicVSR, error) {
	cmd := exec.Command("nvidia-smi", "-x", "-q")
	output, err := cmd.Output()

	if err != nil {
		return nil, err
	}

	decoder := xml2map.NewDecoder(strings.NewReader(string(output)))
	result, err := decoder.Decode()

	if err != nil {
		return nil, err
	}

	driverVersion := result["nvidia_smi_log"].(map[string]interface{})["driver_version"].(string)
	g.DriverVersion = driverVersion

	cudaVersion := result["nvidia_smi_log"].(map[string]interface{})["cuda_version"].(string)
	g.CudaVersion = cudaVersion

	gpuInfo := result["nvidia_smi_log"].(map[string]interface{})["gpu"].(map[string]interface{})
	jsonBody, err := json.Marshal(gpuInfo)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(jsonBody, g); err != nil {
		return nil, err
	}

	//uc.log.WithContext(ctx).Infof("CreateBasicVSR: %v", g)

	return uc.repo.Save(ctx, g)
}

func (uc *BasicVSRUsecase) ExecBasicVsr(ctx context.Context, g *GPURequest) (*ExecResult, error) {

	command := fmt.Sprintf("%s && %s && %s && %s",
		"cd /root/pytorch-env",
		"source ./larry/bin/activate",
		"cd ./RealBasicVSR/fork/RealBasicVSR",
		"nohup python inference_realbasicvsr_fixMem.py configs/realbasicvsr_x4.py checkpoints/RealBasicVSR_x4.pth data/input3.mp4 results/output333.mp4 --fps=30 --max_seq_len=20 --split=3 > /root/pytorch-env/RealBasicVSR/fork/RealBasicVSR/console.log 2>&1 &")

	cmd := exec.Command("/bin/sh", "-c", command)
	output, err := cmd.Output()

	if err != nil {
		return nil, err
	}
	message := string(output)
	uc.log.WithContext(ctx).Infof("ExecBasicVsr: %v", message)

	return &ExecResult{Name: message}, nil
}
