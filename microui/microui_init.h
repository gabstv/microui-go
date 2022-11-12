#include "microui.h"
#include <stdlib.h>


#if defined(__cplusplus)
extern "C" {            // Prevents name mangling of functions
#endif

typedef int (*text_width_fn)(mu_Font font, const char *str, int len);

mu_Context* NewContext(void);

void SetTextWidthFunc(mu_Context* ctx, text_width_fn fn);

#if defined(__cplusplus)
}
#endif