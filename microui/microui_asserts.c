#include "microui.h"
#include "microui_init.h"

int SizeOfCommand(void) {
    mu_Command cmd;
    cmd.rect.base.type = MU_COMMAND_RECT;
    cmd.rect.color = (mu_Color){1, 1, 1, 1};
    cmd.rect.rect = (mu_Rect){1, 1, 1, 1};

    return sizeof(cmd);
}

int SizeOfEnum(void) {
    return sizeof(MU_COMMAND_JUMP);
}
