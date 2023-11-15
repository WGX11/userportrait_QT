#ifndef COLLECTORCONTROLLER_H
#define COLLECTORCONTROLLER_H

#include <QObject>
#include <QDialog>
#include <QPushButton>
#include <QVBoxLayout>
#include <QLabel>

class CollectorController : public QObject
{
    Q_OBJECT

public:
    explicit CollectorController(QObject *parent = nullptr);

    void showControlDialog();

signals:
    void collectionStatusChanged(bool isCollecting);

private:
    QDialog *controlDialog;
    QPushButton *collectButton;
    QLabel *statusLabel;
    bool isCollecting;

private:
    void toggleCollecting();
};

#endif // COLLECTORCONTROLLER_H
