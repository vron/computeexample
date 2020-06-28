// Code generated DO NOT EDIT

#define uint unsigned int

typedef struct {
	float* imgData;
	uint imgWidth;
} kernel_data;

void* cpt_new_kernel(int);
int cpt_dispatch_kernel(void*, kernel_data, int, int, int);
void cpt_free_kernel(void*);
