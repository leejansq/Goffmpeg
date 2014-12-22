// Ffmpeg project Ffmpeg.go
package ffmpeg

/*

#include "stdafx.h"
#include <windows.h>
#include "libavcodec/avcodec.h"
#include "libavutil/avutil.h"
#include "libswscale/swscale.h"
#include "snprintf.h"

#define INBUF_SIZE 1024*10

uint8_t* pFrameBuffer = NULL;


HINSTANCE hDLLB; //定义DLL包柄
HINSTANCE swsDLL;
typedef void ( *func)();  //定义函数指针原型
typedef AVCodec* ( *funv)();
typedef AVCodecContext* ( *funcon)();
typedef int ( *funi)();
typedef AVFrame* ( *funm)();
typedef struct SwsContext* ( *funs)();

AVCodecContext* pCodecContext;
AVFrame* pFrame;
//AVFrame* pFrameRGB;
AVCodec* pCodec;
AVPacket Packet;

func Cavcodec_register_all;
func Cav_init_packet;
func Csws_scale;
funv Cavcodec_find_decoder;
funcon Cavcodec_alloc_context3;
funi Cavcodec_open2;
funi Cavcodec_decode_video2;
funm Cavcodec_alloc_frame;
funs Csws_getCachedContext;

int intDllA()
{
//FreeLibrary(hDLL);
//FreeLibrary(AVhDLL);
 printf("Starting....\n");
 if (hDLLB == NULL)
    hDLLB=LoadLibrary("avcodec-55.dll");  //加载DLL
if (swsDLL == NULL)
    swsDLL=LoadLibrary("swscale-2.dll");  //加载DLL
 Cavcodec_register_all = (func)GetProcAddress(hDLLB,"avcodec_register_all");
 Cavcodec_find_decoder = (funv)GetProcAddress(hDLLB,"avcodec_find_decoder");
 Cavcodec_alloc_context3 = (funcon)GetProcAddress(hDLLB,"avcodec_alloc_context3");
 Cavcodec_open2 = (funi)GetProcAddress(hDLLB,"avcodec_open2");
 Cavcodec_decode_video2 =(funi)GetProcAddress(hDLLB,"avcodec_decode_video2");
 Cavcodec_alloc_frame = (funm)GetProcAddress(hDLLB,"avcodec_alloc_frame");
 Cav_init_packet = (func)GetProcAddress(hDLLB,"av_init_packet");

 Csws_getCachedContext = (funs)GetProcAddress(swsDLL,"sws_getCachedContext");
 Csws_scale = (func)GetProcAddress(swsDLL,"sws_scale");


 printf("goto Init()....\n");
	//Cavcodec_init();
    Cavcodec_register_all();
	pCodec = Cavcodec_find_decoder(CODEC_ID_H264);
	if (NULL == pCodec) {
        fprintf(stderr, "!! find h264 decoder failed\n");
		return -1;
        //exit(1);
    }

    pCodecContext = Cavcodec_alloc_context3(pCodec);
    if (NULL == pCodecContext) {
        fprintf(stderr, "!! Could not allocate video codec context\n");
		return -1;
        //exit(1);
    }
	else{
		//pCodecContext->width=640;
		//pCodecContext->height=360;
	}
	 //if (pCodec->capabilities & CODEC_CAP_TRUNCATED) {
        //pCodecContext->flags |= CODEC_FLAG_TRUNCATED;
		pCodecContext->flags |= CODEC_FLAG_EMU_EDGE | CODEC_FLAG_LOW_DELAY;
		pCodecContext->debug |= FF_DEBUG_MMCO;
    	pCodecContext->pix_fmt = PIX_FMT_YUV420P;
    //}

    if (Cavcodec_open2(pCodecContext, pCodec, NULL) < 0) {
        fprintf(stderr, "!! Could not open codec\n");
		return -1;
        //exit(1);
    }

	 pFrame = Cavcodec_alloc_frame();
    if (NULL == pFrame) {
        fprintf(stderr, "Could not allocate video frame\n");
		return -1;
        //exit(1);
    }

    //AVPacket Packet;
    Cav_init_packet(&Packet);
	return 0;
}

int San(char* buf ,int size)
{
    printf("goto San()....\n");


    Packet.size=size;
	Packet.data=buf;
    int nFrameCount = 0;
	int nGotFrame = 0;
    int nLen = Cavcodec_decode_video2(pCodecContext, pFrame, &nGotFrame, &Packet);
	printf("%d",pFrame->linesize[0]);
	if(nGotFrame>0){
		return 1;
	}
		return 0;



 //       pImgConvertCtx = Csws_getCachedContext(pImgConvertCtx, pCodecContext->width, pCodecContext->height, pCodecContext->pix_fmt,
 //           pCodecContext->width, pCodecContext->height, PIX_FMT_RGB24, SWS_BICUBIC, NULL, NULL, NULL);

 //       Csws_scale(pImgConvertCtx, (const uint8_t* const*)pFrame->data, pFrame->linesize, 0,
 //           pCodecContext->height, pFrameRGB->data, pFrameRGB->linesize);
		//	printf("hi niu %c",pFrameRGB->data[0]);
    //DecodeH264("test.h264", "test_out.yuv");
	 //avcodec_close(pCodecContext);
    //av_free(pCodecContext);
    //avcodec_free_frame(&pFrame);
    //av_free(pFrameBuffer);
    //avcodec_free_frame(&pFrameRGB);
    //return 0;
}
*/
import "C"

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"unsafe"
)

func CArrayToGoArray(cArray unsafe.Pointer, size int) (goArray []uint8) {
	p := uintptr(cArray)
	for i := 0; i < size; i++ {
		j := *(*uint8)(unsafe.Pointer(p))
		goArray = append(goArray, j)
		p += unsafe.Sizeof(j)
	}
	return
}

func Decoder_register() {
	fmt.Println(C.intDllA())
}
func DecoderH264(data unsafe.Pointer, size int) (*image.RGBA, int) {
	ret := C.San((*C.char)(data), C.int(size))
	if int(ret) == 0 {
		return nil, 0
	}
	pFrame := C.pFrame
	fmt.Println(pFrame.linesize)
	w := int(pFrame.width)
	h := int(pFrame.height)
	mgic := image.NewYCbCr(image.Rect(0, 0, w, h), image.YCbCrSubsampleRatio420)
	mgic.Y = CArrayToGoArray(unsafe.Pointer(pFrame.data[0]), int(pFrame.linesize[0])*h)
	mgic.Cb = CArrayToGoArray(unsafe.Pointer(pFrame.data[1]), int(pFrame.linesize[1])*h/2)
	mgic.Cr = CArrayToGoArray(unsafe.Pointer(pFrame.data[2]), int(pFrame.linesize[2])*h/2)
	//mgic.YStride = 640
	//mgic.CStride = 320
	fmt.Println("pFrame.width=", int(pFrame.width), "pFrame.height=", int(pFrame.height))
	fmt.Println("mgic.Y=", int(pFrame.linesize[0]), "mgic.Cb=", int(pFrame.linesize[1]), "mgic.Cr=", int(pFrame.linesize[2]))
	fmt.Println(mgic.Rect)
	dstImg := image.NewRGBA(image.Rect(0, 0, w, h))
	//fmt.Println(mgic.Y[640*360-1])
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			dstImg.Set(i, j, mgic.At(i, j))
			//fmt.Println(i, "-", j, ":", mgic.At(i, j))
		}
	}
	return dstImg, 1
}
func yuv2rgb(y color.YCbCr) color.RGBA {
	r, g, b, a := y.RGBA()
	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
}
func Save2Pic(imge *image.RGBA, name string) error {
	fmt.Println("nimaa><")

	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	err = jpeg.Encode(f, imge, &jpeg.Options{100})

	if err != nil {
		return err
	}
	return nil
}

func AAA() {
	Decoder_register()
	//C.San()
	fmt.Printf("SSS")
}

func init() {
	Decoder_register()
	os.Mkdir("PICture", 0777)
	os.Chdir("PICture")
}
