#include "collectorcontroller.h"

CollectorController::CollectorController(QObject *parent) : QObject(parent), isCollecting(false)
{
    controlDialog = new QDialog();
    collectButton = new QPushButton("Begin collecting", controlDialog);
    statusLabel = new QLabel("Collection not started", controlDialog);

    QVBoxLayout *layout = new QVBoxLayout;
    layout->addWidget(statusLabel);
    layout->addWidget(collectButton);
    controlDialog->setLayout(layout);

    connect(collectButton, &QPushButton::clicked, this, &CollectorController::toggleCollecting);
}

void CollectorController::showControlDialog()
{
    controlDialog->exec();
}

void CollectorController::toggleCollecting()
{
    isCollecting = !isCollecting;
    if (isCollecting)
    {
        collectButton->setText("End collecting");
        statusLabel->setText("Be collecting");
    }
    else
    {
        collectButton->setText("Begin collecting");
        statusLabel->setText("Collection not started");
    }
    emit collectionStatusChanged(isCollecting);
}
