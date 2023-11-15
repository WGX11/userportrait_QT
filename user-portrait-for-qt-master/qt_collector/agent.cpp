#include <QCoreApplication>
#include "agent.h"
#include "event_analyzer.h"
#include <QDebug>
#include <QTimer>
#include <QMessageBox>

using qt_collector::Agent;
using qt_collector::UserEventAnalyzer;

Agent *Agent::gAgent_ = nullptr; // agent 实例
QSet<QString> Agent::componentSet;
QFile Agent::recordFile;

Agent::Agent(QObject *parent)
    : eventAnalyzer_(new UserEventAnalyzer(*this))
{
//    connect(&controller, &CollectorController::collectionStatusChanged, this, &Agent::setCollectionStatus);
//    QTimer::singleShot(0, this, &Agent::showCollectorControl);

    QMessageBox::StandardButton reply;
    reply = QMessageBox::question(nullptr, "Data Collection", "Do you agree to record the operation?" ,QMessageBox::Yes | QMessageBox::No);
    bool isCollecting = (reply == QMessageBox::Yes);
    if (!isCollecting)
    {
        return;
    }
    assert(gAgent_ == nullptr);
    gAgent_ = this;

    QString dirPath = QStandardPaths::writableLocation(QStandardPaths::DesktopLocation) +"/data_" +qAppName() + "/";
    // 创建数据目录
    QDir dir(dirPath);
    if (!dir.exists()) {
        QDir().mkpath(dirPath);
    }

    QString pngPath = dirPath + "/png/";
    QDir pngDir(pngPath);
    if (!pngDir.exists())
    {
        QDir().mkpath(pngPath);
    }

    Agent::initRecordFile();

    // 创建数据文件
    qint64 curTimeOfMS = QDateTime::currentDateTime().toMSecsSinceEpoch();
    QString dataFileName = dirPath + QString::number(curTimeOfMS, 10) + ".csv";
    qDebug() << "dataFileName = " << dataFileName;
    dataFile.setFileName(dataFileName);
    openSuccess = dataFile.open(QFile::WriteOnly | QFile::Append);
    qDebug() << "openSuccess = " << openSuccess;


    // qApp: 应用实例，应用退出
    connect(qApp, SIGNAL(aboutToQuit()), this, SLOT(onAppAboutToQuit()));
    // 用户事件
    connect(eventAnalyzer_, SIGNAL(userEvent(QStringList &)),
            this, SLOT(onUserEvent(QStringList &)));
    // 安装事件过滤器
    QCoreApplication::instance()->installEventFilter(eventAnalyzer_);

    onAppStart();
}

Agent::~Agent()
{
    if (dataFile.exists() && openSuccess) {
        dataFile.close();
    }
    closeRecordFile();

}

void Agent::setAddsValue(const QString &value)
{
    componentSet.insert(value);
    if (recordFile.isOpen())
    {
        QTextStream out(&recordFile);
        out << value << "\n";
    }
    else
    {
        qDebug() << "record.txt has not been opened";
    }
}

bool Agent::setContainsValue(const QString &value)
{
    return componentSet.contains(value);
}

void Agent::initRecordFile()
{
    QString dirPath = QStandardPaths::writableLocation(QStandardPaths::DesktopLocation) +"/data_" +qAppName() + "/png/record.txt";
    recordFile.setFileName(dirPath);
    if (recordFile.open(QIODevice::ReadWrite | QIODevice::Text | QIODevice::Append))
    {
        QTextStream in(&recordFile);
        while (!in.atEnd())
        {
            QString line = in.readLine().trimmed();
            if (!line.isEmpty())
            {
                componentSet.insert(line);
            }
        }
    }
    else
    {
        qDebug() << "Cannot open record.txt file";
    }
}

void Agent::closeRecordFile()
{
    recordFile.close();
}

void Agent::showCollectorControl()
{
    controller.showControlDialog();
}

void Agent::setCollectionStatus(bool status)
{
    isCollecting = status;
}

void Agent::onUserEvent(QStringList &list)
{
    writeData(list);
}

void Agent::onAppAboutToQuit()
{
    qint64 curTimeOfMS = QDateTime::currentDateTime().toMSecsSinceEpoch();
    QStringList quitData = {QString::number(AppQuit, 10), QString::number(curTimeOfMS, 10)};
    writeData(quitData);
}

void Agent::onAppStart()
{
    writeData(FileHeader);

    qint64 curTimeOfMS = QDateTime::currentDateTime().toMSecsSinceEpoch();
    QStringList startData = {QString::number(AppStart, 10), QString::number(curTimeOfMS, 10)};
    writeData(startData);
}

void Agent::writeData(QStringList &list)
{
    if (!dataFile.exists() || !openSuccess) {
        return ;
    }

//    if (!isCollecting){
//        return;
//    }

    if (list.isEmpty()) {
        return ;
    }

    qDebug() << "list=" << list;

    QTextStream output(&dataFile);
    for (int idx = 0; idx < list.length(); idx++) {
        output << ConvertString2CSV(list[idx]);
        output << ((idx < list.length() - 1) ? "," : "\n");
    }
}

